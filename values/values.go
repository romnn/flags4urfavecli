package values

import (
	"fmt"
	"strings"
	"time"
)

// DefaultTimestampFormat ...
var DefaultTimestampFormat = "2006-01-02 15:04:05"

// TimestampValue ...
type TimestampValue struct {
	Format  *string
	Default time.Time
	ts      *time.Time
}

func (ts *TimestampValue) getFormat() string {
	if ts.Format != nil {
		return *ts.Format
	}
	return DefaultTimestampFormat
}

// Set ...
func (ts *TimestampValue) Set(value string) error {
	t, err := time.Parse(ts.getFormat(), value)
	if err != nil {
		return fmt.Errorf("%s cannot be parsed as a timestamp (expected format: %s)", value, ts.getFormat())
	}
	ts.ts = &t
	return nil
}

func (ts TimestampValue) String() string {
	if ts.ts != nil {
		return ts.ts.Format(ts.getFormat())
	}
	return ts.Default.Format(ts.getFormat())
}

// EnumValue ...
type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

// Set ...
func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if strings.ToLower(enum) == strings.ToLower(value) {
			e.selected = strings.ToLower(value)
			return nil
		}
	}
	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}
