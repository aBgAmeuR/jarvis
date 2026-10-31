package proxmox

import (
	"context"
	"errors"
	"net/url"
	"slices"
	"strconv"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/maintenance"
	"jarvis/server/internal/domain/supervision"
	"jarvis/server/internal/source/failure"
)

// Virtual machines are left out.
func (c *Client) GuestResources(ctx context.Context, vmids []int) (map[int]maintenance.Resources, error) {
	inventory, err := c.Inventory(ctx)
	if err != nil {
		return nil, err
	}
	var containers []supervision.Guest
	for _, g := range inventory.Guests {
		if g.Kind == supervision.GuestLXC && slices.Contains(vmids, g.VMID) {
			containers = append(containers, g)
		}
	}

	found := make([]maintenance.Resources, len(containers))
	errs := make([]error, len(containers))
	var wg sync.WaitGroup
	for i, g := range containers {
		wg.Go(func() { found[i], errs[i] = c.lxcResources(ctx, g.Node, g.VMID) })
	}
	wg.Wait()
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	resources := make(map[int]maintenance.Resources, len(containers))
	for i, g := range containers {
		resources[g.VMID] = found[i]
	}
	return resources, nil
}

func (c *Client) lxcResources(ctx context.Context, node string, vmid int) (maintenance.Resources, error) {
	var body struct {
		Data lxcConfig `json:"data"`
	}
	path := "/api2/json/nodes/" + url.PathEscape(node) + "/lxc/" + strconv.Itoa(vmid) + "/config"
	if err := c.fetcher.GetJSON(ctx, path, &body); err != nil {
		return maintenance.Resources{}, failure.Wrap(name, err)
	}
	resources, err := body.Data.resources()
	if err != nil {
		return maintenance.Resources{}, &domain.SourceError{Source: name, Kind: domain.SourceInvalidResponse, Err: err}
	}
	return resources, nil
}
