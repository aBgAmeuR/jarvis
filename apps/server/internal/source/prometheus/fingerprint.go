package prometheus

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// fingerprint is Alertmanager's id for a label set: FNV-1a over sorted names and values, each followed by 0xff.
func fingerprint(labels map[string]string) string {
	names := make([]string, 0, len(labels))
	for n := range labels {
		names = append(names, n)
	}
	sort.Strings(names)
	var data []byte
	for _, n := range names {
		data = append(data, n...)
		data = append(data, 0xff)
		data = append(data, labels[n]...)
		data = append(data, 0xff)
	}
	h := fnv.New64a()
	_, _ = h.Write(data)
	return fmt.Sprintf("%016x", h.Sum64())
}
