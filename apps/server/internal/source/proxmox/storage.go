package proxmox

import (
	"strings"

	"jarvis/server/internal/domain/supervision"
)

func (r resource) storage() supervision.Storage {
	return supervision.Storage{
		Name:       r.Storage,
		Node:       r.Node,
		Type:       r.PluginType,
		Content:    strings.FieldsFunc(r.Content, func(c rune) bool { return c == ',' }),
		Status:     r.Status,
		Shared:     r.Shared == 1,
		UsedBytes:  r.Disk,
		TotalBytes: r.MaxDisk,
	}
}
