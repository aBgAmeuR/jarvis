package daemon

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/abgameur/jarvis/internal/config"
	"github.com/abgameur/jarvis/internal/protocol"
	"github.com/abgameur/jarvis/internal/pve"
	"github.com/abgameur/jarvis/internal/secrets"
)

// validate streams one NDJSON event per layer (FR-10) and stops after the
// first failing layer.
func (s *Server) validate(conn net.Conn) {
	enc := json.NewEncoder(conn)

	loaded := loadConfigAndSecrets(s.ConfigDir, s.secretsPath())
	if !emit(enc, protocol.LayerParse, loaded.parseErrs) {
		return
	}
	if !emit(enc, protocol.LayerRequired, loaded.cfg.CheckRequired()) {
		return
	}
	if !emit(enc, protocol.LayerValues, loaded.cfg.CheckValues()) {
		return
	}

	ref := loaded.cfg.Proxmox.Credentials.Ref
	if !emit(enc, protocol.LayerSecrets, loaded.stub.CheckRef(ref)) {
		return
	}
	if !emit(enc, protocol.LayerFingerprint, loaded.stub.CheckFingerprint(ref)) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	runner := s.PVERunner
	if runner == nil {
		r, err := buildSSHRunner(loaded.cfg, loaded.stub, queryTimeout)
		if err != nil {
			emit(enc, protocol.LayerSSH, []error{err})
			return
		}
		runner = r
	}

	if err := pve.CheckSSH(ctx, runner); err != nil {
		emit(enc, protocol.LayerSSH, []error{err})
		return
	}
	emit(enc, protocol.LayerSSH, nil)

	if err := pve.GuestsExist(ctx, runner, loaded.cfg.Jarvis.VMID, loaded.cfg.Docker.VMID); err != nil {
		emit(enc, protocol.LayerGuests, []error{err})
		return
	}
	emit(enc, protocol.LayerGuests, nil)
}

// emit writes one validate event and reports whether the layer passed.
func emit(enc *json.Encoder, layer string, errs []error) bool {
	ev := protocol.ValidateEvent{Layer: layer, OK: len(errs) == 0}
	if !ev.OK {
		ev.Error = joinErrors(errs)
	}
	if err := enc.Encode(ev); err != nil {
		log.Printf("encode validate event: %v", err)
		return false
	}
	if ev.OK {
		log.Printf("validate layer=%s ok", ev.Layer)
	} else {
		log.Printf("validate layer=%s fail: %s", ev.Layer, ev.Error)
	}
	return ev.OK
}

type loaded struct {
	cfg       *config.Config
	stub      *secrets.Stub
	parseErrs []error
}

// loadConfigAndSecrets parses both files and collects parse errors so the
// parse layer can report them together.
func loadConfigAndSecrets(configDir, secretsPath string) loaded {
	out := loaded{cfg: &config.Config{}}
	cfg, err := config.Load(configDir)
	if err != nil {
		out.parseErrs = append(out.parseErrs, err)
	} else {
		out.cfg = cfg
	}
	stub, err := secrets.LoadStub(secretsPath)
	if err != nil {
		out.parseErrs = append(out.parseErrs, err)
	} else {
		out.stub = stub
	}
	return out
}
