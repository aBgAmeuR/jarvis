package search

type Kind string

const (
	KindGuest   Kind = "guest"
	KindService Kind = "service"
	KindProbe   Kind = "probe"
	KindClient  Kind = "client"
	KindPeer    Kind = "peer"
)

// Hit is one object the ⌘K palette can open.
type Hit struct {
	Kind Kind
	// ID is the id in the object's URL.
	ID   string
	Name string
	// Detail tells apart objects of the same name: a VMID, a guest, a target, an IP.
	Detail string
}
