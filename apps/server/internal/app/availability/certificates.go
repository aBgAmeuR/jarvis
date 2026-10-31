package availability

import "time"

type Certificates struct {
	source CertificateSource
	now    func() time.Time
}

func NewCertificates(source CertificateSource) *Certificates {
	return &Certificates{source: source, now: time.Now}
}
