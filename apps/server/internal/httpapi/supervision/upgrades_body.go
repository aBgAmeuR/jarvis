package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

type upgradesBody struct {
	SecurityPending int  `json:"securityPending" doc:"Pending upgrades from the security repository." example:"3"`
	RegularPending  int  `json:"regularPending" doc:"Other pending upgrades." example:"12"`
	RebootRequired  bool `json:"rebootRequired" doc:"Whether a reboot is needed to finish an upgrade." example:"false"`
}

func newUpgradesBody(u *domainsupervision.Upgrades) *upgradesBody {
	if u == nil {
		return nil
	}
	return &upgradesBody{SecurityPending: u.Security, RegularPending: u.Regular, RebootRequired: u.RebootRequired}
}
