package secrets

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const PlaceholderFingerprint = "SHA256:replace-with-pve-host-key"

var fingerprintRE = regexp.MustCompile(`^SHA256:[A-Za-z0-9+/]{43}$`)

type Material struct {
	PrivateKeyPath     string `yaml:"private_key_path"`
	HostKeyFingerprint string `yaml:"host_key_fingerprint"`
}

type Stub struct {
	entries map[string]Material
}

func LoadStub(path string) (*Stub, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open secrets %q: %w", path, err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)

	var entries map[string]Material
	if err := dec.Decode(&entries); err != nil {
		return nil, fmt.Errorf("parse secrets %q: %w", path, err)
	}
	if entries == nil {
		entries = map[string]Material{}
	}
	return &Stub{entries: entries}, nil
}

func (s *Stub) Resolve(provider, ref string) (Material, error) {
	if provider != "stub" {
		return Material{}, fmt.Errorf("unsupported secrets provider %q", provider)
	}
	m, ok := s.entries[ref]
	if !ok {
		return Material{}, fmt.Errorf("unknown secrets ref %q", ref)
	}
	return m, nil
}

func (s *Stub) CheckRef(ref string) []error {
	var errs []error
	m, ok := s.entries[ref]
	if !ok {
		return []error{fmt.Errorf("secrets ref %q not found", ref)}
	}
	if strings.TrimSpace(m.PrivateKeyPath) == "" {
		errs = append(errs, fmt.Errorf("secrets ref %q: private_key_path is required", ref))
		return errs
	}
	expanded, err := ExpandPath(m.PrivateKeyPath)
	if err != nil {
		errs = append(errs, fmt.Errorf("secrets ref %q: %w", ref, err))
		return errs
	}
	info, err := os.Stat(expanded)
	if err != nil {
		errs = append(errs, fmt.Errorf("secrets ref %q: private key %q: %w", ref, expanded, err))
		return errs
	}
	if info.IsDir() {
		errs = append(errs, fmt.Errorf("secrets ref %q: private key %q is a directory", ref, expanded))
	}
	f, err := os.Open(expanded)
	if err != nil {
		errs = append(errs, fmt.Errorf("secrets ref %q: private key %q unreadable: %w", ref, expanded, err))
		return errs
	}
	_ = f.Close()
	return errs
}

func (s *Stub) CheckFingerprint(ref string) []error {
	m, ok := s.entries[ref]
	if !ok {
		return []error{fmt.Errorf("secrets ref %q not found", ref)}
	}
	fp := strings.TrimSpace(m.HostKeyFingerprint)
	if fp == "" {
		return []error{fmt.Errorf("secrets ref %q: host_key_fingerprint is required", ref)}
	}
	if fp == PlaceholderFingerprint {
		return []error{fmt.Errorf("secrets ref %q: host_key_fingerprint is still the example placeholder", ref)}
	}
	if !fingerprintRE.MatchString(fp) {
		return []error{fmt.Errorf("secrets ref %q: host_key_fingerprint must be SHA256: + 43 unpadded base64 characters", ref)}
	}
	return nil
}

func ExpandPath(path string) (string, error) {
	if path == "~" || strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("expand home in %q: %w", path, err)
		}
		if path == "~" {
			return home, nil
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

func ValidFingerprint(fp string) bool {
	return fingerprintRE.MatchString(strings.TrimSpace(fp))
}
