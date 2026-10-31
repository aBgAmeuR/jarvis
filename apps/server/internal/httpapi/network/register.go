package network

import (
	"github.com/danielgtaylor/huma/v2"

	appnetwork "jarvis/server/internal/app/network"
	"jarvis/server/internal/httpapi/problem"
)

type Services struct {
	Clients *appnetwork.Clients
	DNS     *appnetwork.DNS
	Peers   *appnetwork.Peers
}

func Register(api huma.API, problems *problem.Mapper, s Services) {
	registerListClients(api, problems, s.Clients)
	registerGetClient(api, problems, s.Clients)
	registerListClientDomains(api, problems, s.Clients)
	registerGetClientMetrics(api, problems, s.Clients)
	registerGetDNS(api, problems, s.DNS)
	registerGetDNSMetrics(api, problems, s.DNS)
	registerListDNSDomains(api, problems, s.DNS)
	registerListPeers(api, problems, s.Peers)
	registerGetPeer(api, problems, s.Peers)
	registerGetPeerMetrics(api, problems, s.Peers)
}
