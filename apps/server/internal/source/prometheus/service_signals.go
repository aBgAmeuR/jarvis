package prometheus

import (
	"context"
	"strconv"
	"sync"

	"jarvis/server/internal/domain/services"
	"jarvis/server/internal/source/failure"
)

func (c *Client) ServiceSignals(ctx context.Context) (services.Signals, error) {
	exprs := [...]string{
		`node_systemd_unit_state{state="active"}`,
		`docker_container_ok`,
		`probe_success{job="probe_http"}`,
		`count by (client, server) (traces_service_graph_request_total{source="beyla"})`,
		`count by (service_name) (http_server_request_duration_seconds_count)`,
	}
	var results [len(exprs)][]entry
	var errs [len(exprs)]error
	var wg sync.WaitGroup
	for i, expr := range exprs {
		wg.Go(func() { results[i], errs[i] = c.queryVector(ctx, expr) })
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return services.Signals{}, failure.Wrap(name, err)
		}
	}

	signals := services.Signals{Sites: map[string]bool{}, Instrumented: map[string]bool{}}
	for _, e := range results[0] {
		if vmid, err := strconv.Atoi(e.labels["vmid"]); err == nil {
			signals.Units = append(signals.Units, services.SystemdUnit{VMID: vmid, Unit: e.labels["name"], Active: e.value == 1})
		}
	}
	for _, e := range results[1] {
		if vmid, err := strconv.Atoi(e.labels["vmid"]); err == nil {
			signals.Containers = append(signals.Containers, services.Container{VMID: vmid, Name: e.labels["compose_service"], OK: e.value == 1})
		}
	}
	for _, e := range results[2] {
		signals.Sites[e.labels["site"]] = e.value == 1
	}
	for _, e := range results[3] {
		signals.Edges = append(signals.Edges, services.Edge{Client: e.labels["client"], Server: e.labels["server"]})
	}
	for _, e := range results[4] {
		signals.Instrumented[e.labels["service_name"]] = true
	}
	return signals, nil
}
