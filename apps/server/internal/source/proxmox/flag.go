package proxmox

import "encoding/json"

// flag is a boolean that Proxmox writes as true, false, 1 or 0.
type flag bool

func (f *flag) UnmarshalJSON(data []byte) error {
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*f = flag(b)
		return nil
	}
	var n int
	if err := json.Unmarshal(data, &n); err != nil {
		return err
	}
	*f = n != 0
	return nil
}
