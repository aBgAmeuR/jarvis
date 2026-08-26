// Package pve talks to the Proxmox node: an SSH transport (Runner) and the
// pvesh/pct commands built on top of it.
package pve

import (
	"bytes"
	"context"
	"crypto/subtle"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/abgameur/jarvis/internal/secrets"
)

const defaultTimeout = 15 * time.Second

// Runner runs a remote command and returns stdout.
type Runner interface {
	Run(ctx context.Context, argv []string) ([]byte, error)
}

type SSHConfig struct {
	Address            string
	User               string
	PrivateKeyPath     string
	HostKeyFingerprint string
	Timeout            time.Duration
}

type SSHRunner struct {
	cfg SSHConfig
}

func NewSSHRunner(cfg SSHConfig) (*SSHRunner, error) {
	if cfg.Timeout <= 0 {
		cfg.Timeout = defaultTimeout
	}
	keyPath, err := secrets.ExpandPath(cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	cfg.PrivateKeyPath = keyPath
	return &SSHRunner{cfg: cfg}, nil
}

func (r *SSHRunner) Run(ctx context.Context, argv []string) ([]byte, error) {
	if len(argv) == 0 {
		return nil, fmt.Errorf("empty remote argv")
	}

	key, err := os.ReadFile(r.cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read private key: %w", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	want := strings.TrimSpace(r.cfg.HostKeyFingerprint)
	sshCfg := &ssh.ClientConfig{
		User: r.cfg.User,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: func(_ string, _ net.Addr, key ssh.PublicKey) error {
			got := ssh.FingerprintSHA256(key)
			if subtle.ConstantTimeCompare([]byte(got), []byte(want)) != 1 {
				return fmt.Errorf("ssh: host key fingerprint mismatch: got %s, want %s", got, want)
			}
			return nil
		},
		Timeout: r.cfg.Timeout,
	}

	addr := net.JoinHostPort(r.cfg.Address, "22")
	type dialResult struct {
		client *ssh.Client
		err    error
	}
	ch := make(chan dialResult, 1)
	go func() {
		client, err := ssh.Dial("tcp", addr, sshCfg)
		ch <- dialResult{client: client, err: err}
	}()

	var client *ssh.Client
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("ssh dial %s: %w", addr, ctx.Err())
	case dr := <-ch:
		if dr.err != nil {
			return nil, fmt.Errorf("ssh dial %s: %w", addr, dr.err)
		}
		client = dr.client
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("ssh session: %w", err)
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	cmd := shellJoin(argv)
	done := make(chan error, 1)
	go func() {
		done <- session.Run(cmd)
	}()

	select {
	case <-ctx.Done():
		_ = session.Close()
		return nil, fmt.Errorf("ssh run: %w", ctx.Err())
	case err := <-done:
		if err != nil {
			msg := strings.TrimSpace(stderr.String())
			if msg == "" {
				msg = err.Error()
			}
			return nil, fmt.Errorf("remote %s: %s", argv[0], msg)
		}
		return stdout.Bytes(), nil
	}
}

func shellJoin(argv []string) string {
	parts := make([]string, len(argv))
	for i, a := range argv {
		parts[i] = strconv.Quote(a)
	}
	return strings.Join(parts, " ")
}
