package loki

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// samplePoint reads the [unix seconds, "number"] pair of a range query.
type samplePoint struct {
	time  time.Time
	value float64
}

func (p *samplePoint) UnmarshalJSON(data []byte) error {
	var pair [2]json.RawMessage
	if err := json.Unmarshal(data, &pair); err != nil {
		return fmt.Errorf("sample: %w", err)
	}
	var seconds float64
	if err := json.Unmarshal(pair[0], &seconds); err != nil {
		return fmt.Errorf("sample time: %w", err)
	}
	var text string
	if err := json.Unmarshal(pair[1], &text); err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	p.time = time.UnixMilli(int64(seconds * 1000)).UTC()
	p.value = value
	return nil
}
