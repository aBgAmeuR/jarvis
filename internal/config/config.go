package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services []Service
}

type Service struct {
	Name string `yaml:"name"`
}

func ReadConfig(filePath string) (*Config, error) {
	log.Println("read_config")

	serviceFiles, err := filepath.Glob(filepath.Join(filePath, "services", "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("find service config files: %w", err)
	}

	config := Config{Services: make([]Service, 0, len(serviceFiles))}
	for _, serviceFile := range serviceFiles {
		contents, err := os.ReadFile(serviceFile)
		if err != nil {
			return nil, fmt.Errorf("read service config %q: %w", serviceFile, err)
		}

		var service Service
		if err := yaml.Unmarshal(contents, &service); err != nil {
			return nil, fmt.Errorf("parse service config %q: %w", serviceFile, err)
		}

		config.Services = append(config.Services, service)
	}

	return &config, nil
}
