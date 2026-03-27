package models

import (
	"fmt"
	"strings"
	"time"
)

type Rastamp string
type ChangedBy string
type ChangedDate time.Time
type CreatedBy string
type CreatedDate time.Time

// Returns data within rastamp.
// On error, returned fields will be emty exept error.
func (r Rastamp) Unwrap() (ChangedBy, ChangedDate, CreatedBy, CreatedDate, error) {

	parts := strings.Split(strings.TrimPrefix(string(r), "/"), "/")

	if len(parts) == 4 {
		changedBy := ChangedBy(parts[1])
		createdBy := CreatedBy(parts[3])

		t1, err := time.Parse("20060102150405", parts[0])
		if err != nil {
			return "", ChangedDate{}, "", CreatedDate{}, err
		}
		t2, err := time.Parse("20060102150405", parts[2])
		if err != nil {
			return "", ChangedDate{}, "", CreatedDate{}, err
		}

		ChangedDate := ChangedDate(t1)
		CreatedDate := CreatedDate(t2)

		return changedBy, ChangedDate, createdBy, CreatedDate, nil
	}

	return "", ChangedDate{}, "", CreatedDate{}, fmt.Errorf("insufficent lenght in rastamp, got: %v", parts)
}

// Returns data within rastamp.
// Wrapper over Rastamp.Unwrap(). Will return emty fileds on slient error.
func (r Rastamp) MustUnwrap() (ChangedBy, ChangedDate, CreatedBy, CreatedDate) {

	changedBy, changedDate, createdBy, CreatedDate, _ := r.Unwrap()
	return changedBy, changedDate, createdBy, CreatedDate
}
