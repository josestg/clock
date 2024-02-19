package clock

import (
	"fmt"
	"time"
)

// Clock knows how to capture the current time.
type Clock interface {
	// Now returns the current time.
	Now() time.Time
}

// Func is an adapter to allow the use of ordinary functions as Clocks.
type Func func() time.Time

func (f Func) Now() time.Time { return f() }

// Static returns a Clock that returns the same time every call to Now.
func Static(now time.Time) Clock { return Func(func() time.Time { return now }) }

type withLocation struct {
	loc *time.Location
}

// FromLocation returns a Clock that captures the current time in the given location.
func FromLocation(loc *time.Location) Clock {
	return &withLocation{loc: loc}
}

// LoadLocation creates a Clock that captures the current time in the given location name.
// It returns an error if the location name is invalid.
func LoadLocation(name string) (Clock, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, fmt.Errorf("load location %q: %w", name, err)
	}
	return FromLocation(loc), nil
}

func (c *withLocation) Now() time.Time { return time.Now().In(c.loc) }

type singleton uint8

const (
	Zero singleton = iota
	UTC
	Local
)

var zeroTime time.Time

func (s singleton) Now() time.Time {
	switch s {
	case Zero:
		return zeroTime
	case UTC:
		return time.Now().In(time.UTC)
	case Local:
		return time.Now().In(time.Local)
	default:
		panic("clock: unknown singleton")
	}
}
