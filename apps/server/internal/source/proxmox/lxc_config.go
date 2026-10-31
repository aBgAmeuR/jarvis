package proxmox

import (
	"fmt"
	"strconv"
	"strings"

	"jarvis/server/internal/domain/maintenance"
)

// Proxmox leaves out the memory and the swap of a container that keeps the default of 512 MiB.
const defaultLXCMemoryMB = 512

type lxcConfig struct {
	Cores  int    `json:"cores"`
	Memory *int   `json:"memory"`
	Swap   *int   `json:"swap"`
	Rootfs string `json:"rootfs"`
}

func (c lxcConfig) resources() (maintenance.Resources, error) {
	disk, err := rootfsGB(c.Rootfs)
	if err != nil {
		return maintenance.Resources{}, err
	}
	return maintenance.Resources{
		Cores:    c.Cores,
		MemoryMB: valueOr(c.Memory, defaultLXCMemoryMB),
		SwapMB:   valueOr(c.Swap, defaultLXCMemoryMB),
		DiskGB:   disk,
	}, nil
}

func valueOr(v *int, fallback int) int {
	if v == nil {
		return fallback
	}
	return *v
}

// rootfsGB reads the size of "local-lvm:vm-111-disk-0,size=64G". Without a unit the size is in bytes.
func rootfsGB(rootfs string) (int, error) {
	for option := range strings.SplitSeq(rootfs, ",") {
		size, ok := strings.CutPrefix(option, "size=")
		if !ok || size == "" {
			continue
		}
		factor := int64(1)
		digits := size
		switch size[len(size)-1] {
		case 'K':
			factor = 1 << 10
		case 'M':
			factor = 1 << 20
		case 'G':
			factor = 1 << 30
		case 'T':
			factor = 1 << 40
		}
		if factor > 1 {
			digits = size[:len(size)-1]
		}
		n, err := strconv.ParseInt(digits, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("rootfs size %q: %w", size, err)
		}
		return int(n * factor >> 30), nil
	}
	return 0, fmt.Errorf("rootfs %q has no size", rootfs)
}
