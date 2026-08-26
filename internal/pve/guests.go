package pve

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/abgameur/jarvis/internal/protocol"
)

type nodeInfo struct {
	Node string `json:"node"`
}

type guestInfo struct {
	VMID int `json:"vmid"`
}

type guestListItem struct {
	VMID   int    `json:"vmid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func pveshGet(ctx context.Context, runner Runner, path string) ([]byte, error) {
	return runner.Run(ctx, []string{
		"sudo", "-n", "/usr/bin/pvesh", "get", path, "--output-format", "json",
	})
}

func discoverNode(ctx context.Context, runner Runner) (string, error) {
	nodesRaw, err := pveshGet(ctx, runner, "/nodes")
	if err != nil {
		return "", err
	}
	var nodes []nodeInfo
	if err := json.Unmarshal(nodesRaw, &nodes); err != nil {
		return "", fmt.Errorf("parse /nodes json: %w", err)
	}
	if len(nodes) != 1 {
		return "", fmt.Errorf("expected exactly one Proxmox node, got %d", len(nodes))
	}
	node := nodes[0].Node
	if node == "" {
		return "", fmt.Errorf("proxmox node name is empty")
	}
	return node, nil
}

// GuestsExist checks that jarvisVMID and dockerVMID appear on the node's LXC or QEMU lists.
func GuestsExist(ctx context.Context, runner Runner, jarvisVMID, dockerVMID int) error {
	node, err := discoverNode(ctx, runner)
	if err != nil {
		return err
	}

	vmids := make(map[int]struct{})
	for _, kind := range []string{"lxc", "qemu"} {
		path := fmt.Sprintf("/nodes/%s/%s", node, kind)
		raw, err := pveshGet(ctx, runner, path)
		if err != nil {
			return err
		}
		var guests []guestInfo
		if err := json.Unmarshal(raw, &guests); err != nil {
			return fmt.Errorf("parse %s json: %w", path, err)
		}
		for _, g := range guests {
			vmids[g.VMID] = struct{}{}
		}
	}

	var missing []string
	if _, ok := vmids[jarvisVMID]; !ok {
		missing = append(missing, fmt.Sprintf("jarvis.vmid %d not on node", jarvisVMID))
	}
	if _, ok := vmids[dockerVMID]; !ok {
		missing = append(missing, fmt.Sprintf("docker.vmid %d not on node", dockerVMID))
	}
	if len(missing) > 0 {
		return fmt.Errorf("%s", strings.Join(missing, "; "))
	}
	return nil
}

// ListGuests returns LXC and QEMU guests, sorted by ID. TYPE qemu is mapped to vm.
func ListGuests(ctx context.Context, runner Runner) ([]protocol.Guest, error) {
	node, err := discoverNode(ctx, runner)
	if err != nil {
		return nil, err
	}

	out := make([]protocol.Guest, 0)
	seen := make(map[int]struct{})
	kinds := []struct {
		pathKind string
		cliType  string
	}{
		{"lxc", protocol.GuestTypeLXC},
		{"qemu", protocol.GuestTypeVM},
	}
	for _, k := range kinds {
		path := fmt.Sprintf("/nodes/%s/%s", node, k.pathKind)
		raw, err := pveshGet(ctx, runner, path)
		if err != nil {
			return nil, err
		}
		var items []guestListItem
		if err := json.Unmarshal(raw, &items); err != nil {
			return nil, fmt.Errorf("parse %s json: %w", path, err)
		}
		for _, item := range items {
			if item.VMID <= 0 {
				return nil, fmt.Errorf("guest missing vmid on %s", path)
			}
			if item.Status == "" {
				return nil, fmt.Errorf("guest %d missing status", item.VMID)
			}
			if item.Status != "running" && item.Status != "stopped" {
				return nil, fmt.Errorf("guest %d: unknown status %q", item.VMID, item.Status)
			}
			if _, dup := seen[item.VMID]; dup {
				return nil, fmt.Errorf("duplicate vmid %d", item.VMID)
			}
			seen[item.VMID] = struct{}{}
			name := item.Name
			if strings.ContainsAny(name, "\t\n") {
				return nil, fmt.Errorf("guest %d: name contains tab or newline", item.VMID)
			}
			if name == "" {
				name = "-"
			}
			out = append(out, protocol.Guest{
				ID:    item.VMID,
				Type:  k.cliType,
				Name:  name,
				Power: item.Status,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// LookupGuest returns the guest with the given vmid.
func LookupGuest(ctx context.Context, runner Runner, vmid int) (protocol.Guest, error) {
	guests, err := ListGuests(ctx, runner)
	if err != nil {
		return protocol.Guest{}, err
	}
	for _, g := range guests {
		if g.ID == vmid {
			return g, nil
		}
	}
	return protocol.Guest{}, fmt.Errorf("guest %d not found on node", vmid)
}

// PowerLXC runs pct start or pct stop on an LXC vmid.
func PowerLXC(ctx context.Context, runner Runner, action string, vmid int) error {
	_, err := runner.Run(ctx, []string{
		"sudo", "-n", "/usr/sbin/pct", action, strconv.Itoa(vmid),
	})
	return err
}

// CheckSSH runs a trivial pvesh call to prove connectivity and auth.
func CheckSSH(ctx context.Context, runner Runner) error {
	_, err := pveshGet(ctx, runner, "/nodes")
	return err
}
