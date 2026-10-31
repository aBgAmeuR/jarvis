package availability

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	appavailability "jarvis/server/internal/app/availability"
	"jarvis/server/internal/httpapi/problem"
)

type bucketBody struct {
	Start        time.Time `json:"start" doc:"First instant of the UTC day, RFC 3339." example:"2026-10-30T00:00:00Z"`
	Availability *float64  `json:"availability,omitempty" doc:"Share of successful probes that day, from 0 to 1. Absent without samples." example:"1"`
}

type getProbeAvailabilityInput struct {
	Probe  string `path:"probe" doc:"Probe id." example:"icmp-116"`
	Bucket string `query:"bucket" enum:"1d" default:"1d" doc:"Size of a bucket. Only days are supported."`
}

type getProbeAvailabilityOutput struct {
	Body struct {
		Items []bucketBody `json:"items" nullable:"false" doc:"The last 90 UTC days, then today so far, oldest first."`
	}
}

func registerGetProbeAvailability(api huma.API, problems *problem.Mapper, probes *appavailability.Probes) {
	huma.Register(api, huma.Operation{
		OperationID: "getProbeAvailability",
		Method:      http.MethodGet,
		Path:        "/probes/{probe}/availability",
		Summary:     "Get the daily availability of a probe",
		Tags:        []string{"availability"},
	}, func(ctx context.Context, in *getProbeAvailabilityInput) (*getProbeAvailabilityOutput, error) {
		buckets, err := probes.Availability(ctx, in.Probe)
		if err != nil {
			return nil, problems.Error(ctx, err)
		}
		out := &getProbeAvailabilityOutput{}
		out.Body.Items = make([]bucketBody, 0, len(buckets))
		for _, b := range buckets {
			out.Body.Items = append(out.Body.Items, bucketBody{Start: b.Start, Availability: b.Availability})
		}
		return out, nil
	})
}
