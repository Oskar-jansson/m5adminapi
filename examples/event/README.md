# Events
The SDK exposes concurrency‑safe wrapper around the eventstream endpoint.

### Concepts

* **Eventstream controller** (`services.Eventstream`) represents a subscription. It exposes a buffered channel (`Stream`) for incoming events and
  convenience helpers such as `OnEvent` and `Close`.
* `HasSeenFirstEvent` is an `atomic.Bool` that flips once the stream receives a
  real (non‑ping/err) event. It is used to determine when the server will
  honour `/event` requests. Until the flag is set the SDK will return
  `sdkError.Err == "eventstream not active"` from
  `RequestEventsToStream()`.
* Callbacks registered with `OnEvent` are executed synchronously on the internal
  reader goroutine. **They must not block** – if you need to do expensive work
  spawn new goroutine.
* An optional overflow handle to notify when the
  internal buffer is full. The stream drops events in that case to avoid
  blocking the reader. **Mhey must not block** – if you need to do expensive work
  spawn new goroutine.
* `StartTime`, `time.Time` for when `StartEventstream()` was called.

### Api behavior

#### Events uses two endpoints provided by adminapi:
- `GET /eventstream` Subscribes to live events.
- `GET /event` Sends old events over SSE stream.

#### Behaviors
Calling `GET /eventstream` returns an SSE stream. Current `Raserver` eventbuffer will be dumped into stream, buffer has around 100+ events in it.

Api provides ping events to help ensure that connection is still valid, should also return error when unexpectadly closed.

From calling `GET /eventstream` to reciving first event takes abut `50 seconds`, any calls to `GET /event` before streams successfully starts will be voided.

### Example

```go
alarmOverflow := func(ev *models.Event) {
    // called when the event buffer is full.
    fmt.Println("overflow!", ev)
}

// open a stream,
// opening stream takes around 50 secs. 
controller, sdkErr := conn.Event.StartEventstream(ctx, 500, alarmOverflow, "include=system,alarm")
if sdkErr != nil {
    panic(sdkErr)
}

// shutdown stream, cancels sub-context.
defer controller.Close()

// send requests once the stream is active.
for {
    sdkErr = conn.Event.RequestEventsToStream(ctx, controller, "?startid=1&count=10")
    if sdkErr != nil {
        if sdkErr.Err == "eventstream not active" {
            // retry until HasSeenFirstEvent clears.
            // Allways protect againts infinite loops (do not do as example)!
            continue
        }
        panic(sdkErr)
    }
    break
}

// register a trigger. Is GoRutine safe.
controller.OnEvent(func(ev *models.Event) {
    if ev.System != nil {
        fmt.Println("system msg:", *ev.System.Message)
    }
})

// you can also access channel directly.
go func() {
    for ev := range controller.Stream {
        println(*ev.System.Id)
    }
}()
```

### Notes

* The SDK does not impliment any controll for back-pressure.
* Callbacks and event channel deliveries occur on the same goroutine that
  reads the HTTP stream.