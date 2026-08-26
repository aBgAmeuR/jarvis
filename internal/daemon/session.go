package daemon

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/abgameur/jarvis/internal/config"
	"github.com/abgameur/jarvis/internal/pve"
	"github.com/abgameur/jarvis/internal/secrets"
)

// session bundles what one op needs: validated config and a Proxmox runner.
type session struct {
	cfg    *config.Config
	runner pve.Runner
}

func (s *Server) secretsPath() string {
	if s.SecretsFile != "" {
		return s.SecretsFile
	}
	return filepath.Join(s.ConfigDir, "secrets.yaml")
}

// newSession loads and checks config and secrets fail-fast, then builds the
// SSH runner. validate does the same walk layer by layer to report each step.
func (s *Server) newSession(timeout time.Duration) (*session, error) {
	cfg, err := config.Load(s.ConfigDir)
	if err != nil {
		return nil, err
	}
	if errs := cfg.CheckRequired(); len(errs) > 0 {
		return nil, fmt.Errorf("%s", joinErrors(errs))
	}
	if errs := cfg.CheckValues(); len(errs) > 0 {
		return nil, fmt.Errorf("%s", joinErrors(errs))
	}

	if s.PVERunner != nil {
		return &session{cfg: cfg, runner: s.PVERunner}, nil
	}

	stub, err := secrets.LoadStub(s.secretsPath())
	if err != nil {
		return nil, err
	}
	ref := cfg.Proxmox.Credentials.Ref
	if errs := stub.CheckRef(ref); len(errs) > 0 {
		return nil, fmt.Errorf("%s", joinErrors(errs))
	}
	if errs := stub.CheckFingerprint(ref); len(errs) > 0 {
		return nil, fmt.Errorf("%s", joinErrors(errs))
	}

	runner, err := buildSSHRunner(cfg, stub, timeout)
	if err != nil {
		return nil, err
	}
	return &session{cfg: cfg, runner: runner}, nil
}

func buildSSHRunner(cfg *config.Config, stub *secrets.Stub, timeout time.Duration) (pve.Runner, error) {
	mat, err := stub.Resolve(cfg.Proxmox.Credentials.Provider, cfg.Proxmox.Credentials.Ref)
	if err != nil {
		return nil, err
	}
	return pve.NewSSHRunner(pve.SSHConfig{
		Address:            cfg.Proxmox.Address,
		User:               cfg.Proxmox.User,
		PrivateKeyPath:     mat.PrivateKeyPath,
		HostKeyFingerprint: mat.HostKeyFingerprint,
		Timeout:            timeout,
	})
}
