package catalog

import (
	"context"
	"encoding/json"
	"os"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/source/failure"
)

const name = "catalog"

// File reads catalog.json on every call.
type File struct {
	path string
}

func NewFile(path string) *File { return &File{path: path} }

func (f *File) Name() string { return name }

func (f *File) Ping(context.Context) error {
	_, err := os.Stat(f.path)
	return failure.Wrap(name, err)
}

type document struct {
	Guests []struct {
		VMID        int      `json:"vmid"`
		Name        string   `json:"name"`
		Address     string   `json:"address"`
		Description string   `json:"description"`
		Units       []string `json:"units"`
		Resources   *struct {
			Cores    int `json:"cores"`
			MemoryMB int `json:"memory_mb"`
			SwapMB   int `json:"swap_mb"`
			DiskGB   int `json:"disk_gb"`
		} `json:"resources"`
	} `json:"guests"`
	Sites []struct {
		Host string `json:"host"`
		URL  string `json:"url"`
		VMID *int   `json:"vmid"`
	} `json:"sites"`
}

func (f *File) read() (document, error) {
	raw, err := os.ReadFile(f.path)
	if err != nil {
		return document{}, failure.Wrap(name, err)
	}
	var doc document
	if err := json.Unmarshal(raw, &doc); err != nil {
		return document{}, &domain.SourceError{Source: name, Kind: domain.SourceInvalidResponse, Err: err}
	}
	return doc, nil
}
