package prometheus

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

type sample struct {
	at    time.Time
	value float64
}

func (s *sample) UnmarshalJSON(data []byte) error {
	var pair [2]json.RawMessage
	if err := json.Unmarshal(data, &pair); err != nil {
		return fmt.Errorf("sample: %w", err)
	}
	var seconds float64
	if err := json.Unmarshal(pair[0], &seconds); err != nil {
		return fmt.Errorf("sample timestamp: %w", err)
	}
	var text string
	if err := json.Unmarshal(pair[1], &text); err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	s.at = time.Unix(0, int64(seconds*float64(time.Second))).UTC()
	s.value = value
	return nil
}

func (s sample) finite() bool { return !math.IsNaN(s.value) && !math.IsInf(s.value, 0) }
