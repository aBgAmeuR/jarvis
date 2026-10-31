package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

type nodeBody struct {
	Name   string     `json:"name" doc:"Proxmox node name." example:"pve"`
	Online bool       `json:"online" doc:"Whether Proxmox reports the node online." example:"true"`
	Usage  *usageBody `json:"usage,omitempty" doc:"Resource usage. Absent when the node is offline or Proxmox gives no figures."`
}

func newNodeBody(n domainsupervision.Node) nodeBody {
	return nodeBody{Name: n.Name, Online: n.Online, Usage: newUsageBody(n.Usage)}
}
