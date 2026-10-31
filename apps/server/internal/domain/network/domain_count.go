package network

type DomainStatus string

const (
	DomainAllowed DomainStatus = "allowed"
	DomainBlocked DomainStatus = "blocked"
)

type DomainCount struct {
	Domain string
	Count  int64
}
