package util

import (
	"strconv"
	"time"
)

// Timestamp is like time.Time, but knows how to unmarshal from JSON Unix timestamp
// numbers and marshal back into the same JSON representation.
type Timestamp struct {
	time.Time
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	sec := float64(t.UnixNano()) * float64(time.Nanosecond) / float64(time.Second)
	return strconv.AppendFloat(nil, sec, 'f', -1, 64), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(0, int64(f*float64(time.Second/time.Nanosecond)))
	return nil
}
