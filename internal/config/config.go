package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Proxmox Proxmox `yaml:"proxmox"`
	Jarvis  Jarvis  `yaml:"jarvis"`
	Docker  Docker  `yaml:"docker"`
}

type Proxmox struct {
	Address     string      `yaml:"address"`
	User        string      `yaml:"user"`
	Credentials Credentials `yaml:"credentials"`
}

type Credentials struct {
	Provider string `yaml:"provider"`
	Ref      string `yaml:"ref"`
}

type Jarvis struct {
	VMID int `yaml:"vmid"`
}

type Docker struct {
	VMID int `yaml:"vmid"`
}

func Load(configDir string) (*Config, error) {
	path := filepath.Join(configDir, "config.yaml")
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config %q: %w", path, err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)

	var cfg Config
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}
	return &cfg, nil
}

func (c *Config) CheckRequired() []error {
	var errs []error
	if strings.TrimSpace(c.Proxmox.Address) == "" {
		errs = append(errs, fmt.Errorf("proxmox.address is required"))
	}
	if strings.TrimSpace(c.Proxmox.User) == "" {
		errs = append(errs, fmt.Errorf("proxmox.user is required"))
	}
	if strings.TrimSpace(c.Proxmox.Credentials.Provider) == "" {
		errs = append(errs, fmt.Errorf("proxmox.credentials.provider is required"))
	}
	if strings.TrimSpace(c.Proxmox.Credentials.Ref) == "" {
		errs = append(errs, fmt.Errorf("proxmox.credentials.ref is required"))
	}
	if c.Jarvis.VMID == 0 {
		errs = append(errs, fmt.Errorf("jarvis.vmid is required"))
	}
	if c.Docker.VMID == 0 {
		errs = append(errs, fmt.Errorf("docker.vmid is required"))
	}
	return errs
}

func (c *Config) CheckValues() []error {
	var errs []error
	if c.Proxmox.Credentials.Provider != "" && c.Proxmox.Credentials.Provider != "stub" {
		errs = append(errs, fmt.Errorf("proxmox.credentials.provider must be %q, got %q", "stub", c.Proxmox.Credentials.Provider))
	}
	if c.Jarvis.VMID != 0 && c.Jarvis.VMID < 100 {
		errs = append(errs, fmt.Errorf("jarvis.vmid must be >= 100, got %d", c.Jarvis.VMID))
	}
	if c.Docker.VMID != 0 && c.Docker.VMID < 100 {
		errs = append(errs, fmt.Errorf("docker.vmid must be >= 100, got %d", c.Docker.VMID))
	}
	if c.Jarvis.VMID != 0 && c.Docker.VMID != 0 && c.Jarvis.VMID == c.Docker.VMID {
		errs = append(errs, fmt.Errorf("docker.vmid must differ from jarvis.vmid (%d)", c.Jarvis.VMID))
	}
	return errs
}
