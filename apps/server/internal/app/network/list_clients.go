package network

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
)

type ClientList struct {
	Items    []network.Client
	Degraded []domain.Degraded
}

// ClientSort names a sort field, prefixed with "-" for descending order.
type ClientSort string

// The leases are required; a failing resolver or Loki only degrades the counts and the names.
func (c *Clients) List(ctx context.Context, sortBy ClientSort, limit int) (ClientList, error) {
	var (
		leases []network.Lease
		usage  map[string]network.Usage
		names  map[string]string

		leasesErr, usageErr, namesErr error
	)
	now := c.now()
	w, err := domain.NewWindow(now.Add(-domain.DefaultWindow), now, now)
	if err != nil {
		return ClientList{}, err
	}
	var wg sync.WaitGroup
	wg.Go(func() { leases, leasesErr = c.source.Leases(ctx) })
	wg.Go(func() { usage, usageErr = c.source.ClientUsage(ctx) })
	wg.Go(func() { names, namesErr = c.names.ClientNames(ctx, w) })
	wg.Wait()
	if leasesErr != nil {
		return ClientList{}, fmt.Errorf("list clients: %w", leasesErr)
	}

	var degraded []domain.Degraded
	if usageErr != nil {
		degraded = append(degraded, domain.NewDegraded("prometheus", usageErr))
	}
	if namesErr != nil {
		degraded = append(degraded, domain.NewDegraded("loki", namesErr))
	}
	clients := join(leases, names, usage)
	if err := sortClients(clients, sortBy); err != nil {
		return ClientList{}, err
	}
	if limit > 0 && len(clients) > limit {
		clients = clients[:limit]
	}
	return ClientList{Items: clients, Degraded: degraded}, nil
}

func sortClients(clients []network.Client, sortBy ClientSort) error {
	field, descending := strings.CutPrefix(string(sortBy), "-")
	var compare func(a, b network.Client) int
	switch field {
	case "", "name":
		compare = func(a, b network.Client) int { return cmp.Compare(displayName(a), displayName(b)) }
	case "queries":
		compare = func(a, b network.Client) int { return cmp.Compare(a.Queries, b.Queries) }
	case "blocked":
		compare = func(a, b network.Client) int { return cmp.Compare(a.Blocked, b.Blocked) }
	case "lastQueryAt":
		compare = compareTimes
	default:
		return fmt.Errorf("%w: cannot sort by %q", domain.ErrInvalid, field)
	}
	slices.SortStableFunc(clients, func(a, b network.Client) int {
		if descending {
			return compare(b, a)
		}
		return compare(a, b)
	})
	return nil
}

func displayName(c network.Client) string {
	if c.Name != "" {
		return strings.ToLower(c.Name)
	}
	return c.IP
}

// compareTimes puts the clients never seen first in ascending order.
func compareTimes(a, b network.Client) int {
	switch {
	case a.LastQueryAt == nil && b.LastQueryAt == nil:
		return 0
	case a.LastQueryAt == nil:
		return -1
	case b.LastQueryAt == nil:
		return 1
	}
	return a.LastQueryAt.Compare(*b.LastQueryAt)
}
