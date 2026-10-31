package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

type storageBody struct {
	Name       string   `json:"name" doc:"Storage id." example:"hdd"`
	Type       string   `json:"type" doc:"Proxmox storage type." example:"zfspool"`
	Content    []string `json:"content" nullable:"false" doc:"What the storage holds." example:"[\"images\",\"rootdir\"]"`
	Status     string   `json:"status" doc:"Proxmox status." example:"available"`
	Shared     bool     `json:"shared" doc:"Whether several nodes use the storage." example:"false"`
	UsedBytes  uint64   `json:"usedBytes" doc:"Space in use." example:"348728942592"`
	TotalBytes uint64   `json:"totalBytes" doc:"Space available." example:"3861712470016"`
}

func newStorageBody(s domainsupervision.Storage) storageBody {
	content := s.Content
	if content == nil {
		content = []string{}
	}
	return storageBody{
		Name:       s.Name,
		Type:       s.Type,
		Content:    content,
		Status:     s.Status,
		Shared:     s.Shared,
		UsedBytes:  s.UsedBytes,
		TotalBytes: s.TotalBytes,
	}
}
