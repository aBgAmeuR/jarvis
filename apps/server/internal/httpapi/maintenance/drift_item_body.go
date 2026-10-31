package maintenance

import domainmaintenance "jarvis/server/internal/domain/maintenance"

type resourceDifferenceBody struct {
	Field    string `json:"field" doc:"Resource that differs." enum:"cores,memoryMb,swapMb,diskGb" example:"memoryMb"`
	Declared int    `json:"declared" doc:"Value in homelab-iac." example:"512"`
	Actual   int    `json:"actual" doc:"Value in Proxmox." example:"1024"`
}

type driftItemBody struct {
	Kind        string                   `json:"kind" doc:"What differs: a guest in Proxmox that homelab-iac does not manage, the reverse, or resources that differ between the two." enum:"missingInIac,missingInProxmox,resourceMismatch" example:"missingInIac"`
	VMID        int                      `json:"vmid" doc:"Guest id." example:"105"`
	Name        string                   `json:"name" doc:"Guest name." example:"remote-dev"`
	Differences []resourceDifferenceBody `json:"differences,omitempty" nullable:"false" doc:"The resources that differ. Only for resourceMismatch."`
}

func newDriftItemBody(i domainmaintenance.DriftItem) driftItemBody {
	body := driftItemBody{Kind: string(i.Kind), VMID: i.Guest.VMID, Name: i.Guest.Name}
	for _, d := range i.Differences {
		body.Differences = append(body.Differences, resourceDifferenceBody{Field: string(d.Field), Declared: d.Declared, Actual: d.Actual})
	}
	return body
}
