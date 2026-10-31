package bootstrap

import (
	appnetwork "jarvis/server/internal/app/network"
	httpnetwork "jarvis/server/internal/httpapi/network"
)

func newNetwork(s sources) httpnetwork.Services {
	return httpnetwork.Services{
		Clients: appnetwork.NewClients(s.prometheus, s.loki, s.loki),
		DNS:     appnetwork.NewDNS(s.prometheus, s.loki),
		Peers:   appnetwork.NewPeers(s.prometheus),
	}
}
