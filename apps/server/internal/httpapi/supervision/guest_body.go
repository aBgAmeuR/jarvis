package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

// GuestBody is exported so GuestDetailBody can embed it in the OpenAPI schema.
type GuestBody struct {
	VMID          int           `json:"vmid" doc:"Proxmox guest id." example:"116"`
	Name          string        `json:"name" doc:"Guest name." example:"dns"`
	Kind          string        `json:"kind" doc:"Container or virtual machine." enum:"lxc,qemu" example:"lxc"`
	Node          string        `json:"node" doc:"Proxmox node hosting the guest." example:"pve"`
	Status        string        `json:"status" doc:"Proxmox status." example:"running"`
	Address       string        `json:"address,omitempty" doc:"IP address from the catalog." example:"192.168.1.116"`
	Description   string        `json:"description,omitempty" doc:"Description from the catalog." example:"Blocky and dnsmasq"`
	ManagedByIac  *bool         `json:"managedByIac,omitempty" doc:"Whether homelab-iac manages the guest. Absent when the catalog is unavailable." example:"true"`
	AlertSeverity string        `json:"alertSeverity,omitempty" doc:"Severity of the worst firing alert on the guest. Absent when none fires." enum:"critical,warning,info" example:"warning"`
	AlertCount    int           `json:"alertCount" doc:"Number of firing alerts on the guest." example:"1"`
	Usage         *usageBody    `json:"usage,omitempty" doc:"Resource usage. Absent unless the guest is running."`
	Upgrades      *upgradesBody `json:"upgrades,omitempty" doc:"Pending apt upgrades. Absent when the guest reports none."`
}

func newGuestBody(g domainsupervision.Guest) GuestBody {
	return GuestBody{
		VMID:          g.VMID,
		Name:          g.Name,
		Kind:          string(g.Kind),
		Node:          g.Node,
		Status:        g.Status,
		Address:       g.Address,
		Description:   g.Description,
		ManagedByIac:  g.ManagedByIaC,
		AlertSeverity: string(g.AlertSeverity),
		AlertCount:    g.AlertCount,
		Usage:         newUsageBody(g.Usage),
		Upgrades:      newUpgradesBody(g.Upgrades),
	}
}
