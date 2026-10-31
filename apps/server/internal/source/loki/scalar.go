package loki

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// sampleValue reads the [timestamp, "number"] pair of an instant query.
type sampleValue float64

func (v *sampleValue) UnmarshalJSON(data []byte) error {
	var pair [2]json.RawMessage
	if err := json.Unmarshal(data, &pair); err != nil {
		return fmt.Errorf("sample: %w", err)
	}
	var text string
	if err := json.Unmarshal(pair[1], &text); err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return fmt.Errorf("sample value: %w", err)
	}
	*v = sampleValue(number)
	return nil
}
