package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
	"sync"
	"sync/atomic"
	"time"
)

type EventService struct {
	client *client.Client
}

func NewEventService(c *client.Client) *EventService {
	return &EventService{client: c}
}

var (
	eventstreamsMu sync.RWMutex
	Eventstreams   = make(map[string]*Eventstream)
)

// controller type for eventstream
type Eventstream struct {
	Stream chan *models.Event

	// turns true when we have seen our first non‑ping event
	HasSeenFirstEvent atomic.Bool

	// user callbacks; they execute synchronously on the reader goroutine.
	Trigger []func(*models.Event)

	// callback when overflow occurs. Should be optimized as to not block reader.
	OverflowTrigger func(*models.Event)

	// cancel function returned from context.WithCancel in SubscribeToStream
	cancel context.CancelFunc

	// meta data
	StartTime time.Time

	mu sync.Mutex
}

// trigger on event
func (es *Eventstream) OnEvent(fn func(*models.Event)) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.Trigger = append(es.Trigger, fn)
}

func (es *Eventstream) Close() {
	es.mu.Lock()
	if es.cancel != nil {
		es.cancel()
	}
	es.mu.Unlock()
}

// internal
func (es *Eventstream) appendToGlobalList(name string) {
	es.mu.Lock()
	defer es.mu.Unlock()
	Eventstreams[name] = es
}

// Starts new eventstream and returns eventstream controller.
//
// Stream takes about 50 sec to open correctly
func (s *EventService) StartEventstream(mainCtx context.Context, bufferSize int, onEventOverflow func(*models.Event), opt ...string) (*Eventstream, *models.SdkError) {

	if bufferSize <= 0 {
		return nil, &models.SdkError{Err: fmt.Errorf("can not open event channel with no buffer size")}
	}
	if len(opt) == 0 {
		return nil, &models.SdkError{Err: fmt.Errorf("must provide options")}
	}

	ctx, cancel := context.WithCancel(mainCtx)
	url := fmt.Sprintf("%s%s/eventstream", s.client.Address, s.client.Path)

	scanner, scannerClose, sdkErr := s.client.RequestStream(ctx, options.QueryParamMerge(url, opt))
	if sdkErr != nil {
		cancel()
		return nil, sdkErr
	}

	es := &Eventstream{
		Stream:          make(chan *models.Event, bufferSize),
		cancel:          cancel,
		OverflowTrigger: onEventOverflow,

		StartTime: time.Now(),
	}
	es.appendToGlobalList(s.client.Credentials.System)

	es.OnEvent(func(ev *models.Event) {
		if !es.HasSeenFirstEvent.Load() && ev.Ping == nil && ev.Error == nil && ev.InternalError == nil {
			es.HasSeenFirstEvent.Store(true)
		}
	})

	// Read and active triggers on event
	go func() {
		s.client.ReadStream(ctx, scanner, func(ev *models.Event) {

			es.mu.Lock()
			for _, t := range es.Trigger {

				t(ev)
			}
			es.mu.Unlock()

			select {
			case es.Stream <- ev:
			default:
				if es.OverflowTrigger != nil {
					es.OverflowTrigger(ev)
				}
				// drops event on overflow to not block.
			}
		}, scannerClose)

		close(es.Stream)
	}()

	return es, nil
}

// Requests events that matches provided options to be sent over eventstream.
//
// Any request made before stream is active will return sdkError.Err = "eventstream not active".
// To avoid complicatd logic to request events, wrap this function in rety
func (s *EventService) RequestEventsToStream(ctx context.Context, es *Eventstream, opt ...string) *models.SdkError {

	if !es.HasSeenFirstEvent.Load() {

		// Any request to /events will be voided if stream is not active.
		return &models.SdkError{Err: fmt.Errorf("eventstream not active")}
	}

	url := fmt.Sprintf("%s/event", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)

}
