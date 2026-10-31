package proxmox

import "jarvis/server/internal/domain/supervision"

type resource struct {
	Type     string  `json:"type"`
	Node     string  `json:"node"`
	Name     string  `json:"name"`
	VMID     int     `json:"vmid"`
	Status   string  `json:"status"`
	Template int     `json:"template"`
	CPU      float64 `json:"cpu"`
	MaxCPU   float64 `json:"maxcpu"`
	Mem      uint64  `json:"mem"`
	MaxMem   uint64  `json:"maxmem"`
	Disk     uint64  `json:"disk"`
	MaxDisk  uint64  `json:"maxdisk"`
	Uptime   uint64  `json:"uptime"`

	Storage    string `json:"storage"`
	PluginType string `json:"plugintype"`
	Content    string `json:"content"`
	Shared     int    `json:"shared"`
}

// usage is nil when Proxmox gives no figures, as for a token without Sys.Audit.
func (r resource) usage() *supervision.Usage {
	if r.MaxMem == 0 {
		return nil
	}
	return &supervision.Usage{
		CPUUsageRatio:    r.CPU,
		CPUCount:         int(r.MaxCPU),
		MemoryUsedBytes:  r.Mem,
		MemoryTotalBytes: r.MaxMem,
		DiskUsedBytes:    r.Disk,
		DiskTotalBytes:   r.MaxDisk,
		UptimeSeconds:    r.Uptime,
	}
}

func (r resource) node() supervision.Node {
	node := supervision.Node{Name: r.Node, Online: r.Status == "online"}
	if node.Online {
		node.Usage = r.usage()
	}
	return node
}

func (r resource) guest() supervision.Guest {
	guest := supervision.Guest{
		VMID:   r.VMID,
		Name:   r.Name,
		Kind:   supervision.GuestKind(r.Type),
		Node:   r.Node,
		Status: r.Status,
	}
	if r.Status == "running" {
		guest.Usage = r.usage()
	}
	return guest
}
