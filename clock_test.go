package clock

import (
	"testing"
	"time"
)

func TestFromLocation(t *testing.T) {
	c := FromLocation(time.UTC)
	got := c.Now()
	if got.Location() != time.UTC {
		t.Errorf("Now() = %v; want %v", got.Location(), time.UTC)
	}
}

func TestLoadLocation(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c, err := LoadLocation("UTC")
		if err != nil {
			t.Fatalf("LoadLocation() error = %v; want nil", err)
		}
		got := c.Now()
		if got.Location() != time.UTC {
			t.Errorf("Now() = %v; want %v", got.Location(), time.UTC)
		}
	})

	t.Run("error", func(t *testing.T) {
		_, err := LoadLocation("unknown")
		if err == nil {
			t.Fatalf("LoadLocation() error = nil; want not nil")
		}
	})
}

func TestFunc_Now(t *testing.T) {
	now := time.Now()
	f := Func(func() time.Time { return now })
	got := f.Now()
	if got != now {
		t.Errorf("Now() = %v; want %v", got, now)
	}
}

func TestSingleton_Now(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		got := Zero.Now()
		if !got.IsZero() {
			t.Errorf("Now() = %v; want zero time", got)
		}
	})

	t.Run("UTC", func(t *testing.T) {
		got := UTC.Now()
		if got.Location() != time.UTC {
			t.Errorf("Now() = %v; want %v", got.Location(), time.UTC)
		}
	})

	t.Run("Local", func(t *testing.T) {
		got := Local.Now()
		if got.Location() != time.Local {
			t.Errorf("Now() = %v; want %v", got.Location(), time.Local)
		}
	})

	t.Run("unknown", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatalf("Now() did not panic")
			}
		}()
		unknown := singleton(0xFF)
		unknown.Now()
	})
}

func TestStatic_Now(t *testing.T) {
	now := time.Now()
	c := Static(now)
	x, y := c.Now(), c.Now()
	if x != now || y != now {
		t.Errorf("Now() = %v, %v; want %v, %v", x, y, now, now)
	}
}
