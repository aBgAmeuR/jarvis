package network

import "time"

type Clients struct {
	source  ClientSource
	names   NameSource
	domains DomainSource
	now     func() time.Time
}

func NewClients(source ClientSource, names NameSource, domains DomainSource) *Clients {
	return &Clients{source: source, names: names, domains: domains, now: time.Now}
}
