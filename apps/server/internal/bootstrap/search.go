package bootstrap

import (
	"context"
	"strconv"

	appavailability "jarvis/server/internal/app/availability"
	appnetwork "jarvis/server/internal/app/network"
	appsearch "jarvis/server/internal/app/search"
	appservices "jarvis/server/internal/app/services"
	"jarvis/server/internal/domain/search"
	"jarvis/server/internal/source/proxmox"
)

func newSearch(s sources, services *appservices.Services, probes *appavailability.Probes, clients *appnetwork.Clients, peers *appnetwork.Peers) *appsearch.Service {
	return appsearch.NewService(
		appsearch.Index{Source: "proxmox", Hits: searchGuests{proxmox: s.proxmox}},
		appsearch.Index{Source: "catalog", Hits: searchServices{services: services}},
		appsearch.Index{Source: "prometheus", Hits: searchProbes{probes: probes}},
		appsearch.Index{Source: "prometheus", Hits: searchClients{clients: clients}},
		appsearch.Index{Source: "prometheus", Hits: searchPeers{peers: peers}},
	)
}

type searchGuests struct {
	proxmox *proxmox.Client
}

func (g searchGuests) SearchHits(ctx context.Context) ([]search.Hit, error) {
	inventory, err := g.proxmox.Inventory(ctx)
	if err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(inventory.Guests))
	for _, guest := range inventory.Guests {
		vmid := strconv.Itoa(guest.VMID)
		hits = append(hits, search.Hit{Kind: search.KindGuest, ID: vmid, Name: guest.Name, Detail: vmid})
	}
	return hits, nil
}

type searchServices struct {
	services *appservices.Services
}

func (s searchServices) SearchHits(ctx context.Context) ([]search.Hit, error) {
	list, err := s.services.List(ctx)
	if err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(list))
	for _, service := range list {
		hits = append(hits, search.Hit{Kind: search.KindService, ID: service.Slug, Name: service.Name, Detail: service.GuestName})
	}
	return hits, nil
}

type searchProbes struct {
	probes *appavailability.Probes
}

func (p searchProbes) SearchHits(ctx context.Context) ([]search.Hit, error) {
	list, err := p.probes.List(ctx)
	if err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(list))
	for _, probe := range list {
		hits = append(hits, search.Hit{Kind: search.KindProbe, ID: probe.ID, Name: probe.Name, Detail: probe.Target})
	}
	return hits, nil
}

type searchClients struct {
	clients *appnetwork.Clients
}

// A client without a name goes by its IP, and its MAC tells it apart.
func (c searchClients) SearchHits(ctx context.Context) ([]search.Hit, error) {
	list, err := c.clients.List(ctx, "", 0)
	if err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(list.Items))
	for _, client := range list.Items {
		hit := search.Hit{Kind: search.KindClient, ID: client.ID, Name: client.Name, Detail: client.IP}
		if hit.Name == "" {
			hit.Name, hit.Detail = client.IP, client.MAC
		}
		hits = append(hits, hit)
	}
	return hits, nil
}

type searchPeers struct {
	peers *appnetwork.Peers
}

func (p searchPeers) SearchHits(ctx context.Context) ([]search.Hit, error) {
	list, err := p.peers.List(ctx)
	if err != nil {
		return nil, err
	}
	hits := make([]search.Hit, 0, len(list))
	for _, peer := range list {
		hits = append(hits, search.Hit{Kind: search.KindPeer, ID: peer.Name, Name: peer.Name, Detail: peer.AllowedIPs})
	}
	return hits, nil
}
