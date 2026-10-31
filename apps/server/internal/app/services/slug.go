package services

import (
	"strconv"

	"jarvis/server/internal/domain/services"
)

// A slug is the name, plus the guest when several share it, plus the type if still ambiguous.
func assignSlugs(list []services.Service) {
	count := map[string]int{}
	for _, s := range list {
		count[s.Name]++
	}
	for i := range list {
		list[i].Slug = list[i].Name
		if count[list[i].Name] > 1 {
			list[i].Slug += "-" + strconv.Itoa(list[i].VMID)
		}
	}
	taken := map[string]int{}
	for _, s := range list {
		taken[s.Slug]++
	}
	for i := range list {
		if taken[list[i].Slug] > 1 {
			list[i].Slug += "-" + string(list[i].Type)
		}
	}
}
