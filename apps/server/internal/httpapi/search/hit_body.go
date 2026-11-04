package search

import "jarvis/server/internal/domain/search"

type SearchHitBody struct {
	Kind   string `json:"kind" enum:"guest,service,probe,client,peer" doc:"Kind of object." example:"service"`
	ID     string `json:"id" doc:"Id of the object, as in its URL." example:"grafana"`
	Name   string `json:"name" doc:"Name of the object." example:"grafana"`
	Detail string `json:"detail" doc:"What tells it apart from objects of the same name: a VMID, a guest, a target, an IP. May be empty." example:"monitoring"`
}

func newSearchHitBody(h search.Hit) SearchHitBody {
	return SearchHitBody{Kind: string(h.Kind), ID: h.ID, Name: h.Name, Detail: h.Detail}
}
