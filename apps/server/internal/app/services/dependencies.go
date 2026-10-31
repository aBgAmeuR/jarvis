package services

import (
	"slices"

	"jarvis/server/internal/domain/services"
)

// Edges name processes, so sites and names shared by several services are left out.
func link(list []services.Service, signals services.Signals) {
	slugOf := map[string]string{}
	ambiguous := map[string]bool{}
	for _, s := range list {
		if s.Type == services.TypeSite {
			continue
		}
		if _, seen := slugOf[s.Name]; seen {
			ambiguous[s.Name] = true
		}
		slugOf[s.Name] = s.Slug
	}
	for i := range list {
		s := &list[i]
		s.Instrumented = signals.Instrumented[s.Name]
		if s.Type == services.TypeSite || ambiguous[s.Name] {
			continue
		}
		for _, edge := range signals.Edges {
			server, known := slugOf[edge.Server]
			if edge.Client == s.Name && known && !ambiguous[edge.Server] && server != s.Slug && !slices.Contains(s.DependsOn, server) {
				s.DependsOn = append(s.DependsOn, server)
			}
		}
		slices.Sort(s.DependsOn)
	}
}
