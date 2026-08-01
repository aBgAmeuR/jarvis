package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services []Service
}

type Service struct {
	Name string `yaml:"name"`
}

func Load(configDir string) (*Config, error) {
	pattern := filepath.Join(configDir, "services", "*.yaml")
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("find service config files: %w", err)
	}

	cfg := &Config{Services: make([]Service, 0, len(paths))}
	for _, path := range paths {
		svc, err := loadServiceFile(path)
		if err != nil {
			return nil, err
		}
		cfg.Services = append(cfg.Services, svc)
	}

	sort.Slice(cfg.Services, func(i, j int) bool {
		return cfg.Services[i].Name < cfg.Services[j].Name
	})

	return cfg, nil
}

func loadServiceFile(path string) (Service, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return Service{}, fmt.Errorf("read service config %q: %w", path, err)
	}

	var svc Service
	if err := yaml.Unmarshal(contents, &svc); err != nil {
		return Service{}, fmt.Errorf("parse service config %q: %w", path, err)
	}

	if err := svc.validate(); err != nil {
		return Service{}, fmt.Errorf("invalid service config %q: %w", path, err)
	}

	return svc, nil
}

func (s Service) validate() error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
