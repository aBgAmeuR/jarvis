package pve

import (
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"
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

type nodeInfo struct {
	Node string `json:"node"`
}

type guestInfo struct {
	VMID int `json:"vmid"`
}

type guestListItem struct {
	VMID   int    `json:"vmid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Guest is a CLI-ready Proxmox row (type lxc/vm, empty name is "-").
type Guest struct {
	ID    int
	Type  string
	Name  string
	Power string
}

func pveshGet(ctx context.Context, runner Runner, path string) ([]byte, error) {
	return runner.Run(ctx, []string{
		"sudo", "-n", "/usr/bin/pvesh", "get", path, "--output-format", "json",
	})
}

func discoverNode(ctx context.Context, runner Runner) (string, error) {
	nodesRaw, err := pveshGet(ctx, runner, "/nodes")
	if err != nil {
		return "", err
	}
	var nodes []nodeInfo
	if err := json.Unmarshal(nodesRaw, &nodes); err != nil {
		return "", fmt.Errorf("parse /nodes json: %w", err)
	}
	if len(nodes) != 1 {
		return "", fmt.Errorf("expected exactly one Proxmox node, got %d", len(nodes))
	}
	node := nodes[0].Node
	if node == "" {
		return "", fmt.Errorf("proxmox node name is empty")
	}
	return node, nil
}

// GuestsExist checks that jarvisVMID and dockerVMID appear on the node's LXC or QEMU lists.
func GuestsExist(ctx context.Context, runner Runner, jarvisVMID, dockerVMID int) error {
	node, err := discoverNode(ctx, runner)
	if err != nil {
		return err
	}

	vmids := make(map[int]struct{})
	for _, kind := range []string{"lxc", "qemu"} {
		path := fmt.Sprintf("/nodes/%s/%s", node, kind)
		raw, err := pveshGet(ctx, runner, path)
		if err != nil {
			return err
		}
		var guests []guestInfo
		if err := json.Unmarshal(raw, &guests); err != nil {
			return fmt.Errorf("parse %s json: %w", path, err)
		}
		for _, g := range guests {
			vmids[g.VMID] = struct{}{}
		}
	}

	var missing []string
	if _, ok := vmids[jarvisVMID]; !ok {
		missing = append(missing, fmt.Sprintf("jarvis.vmid %d not on node", jarvisVMID))
	}
	if _, ok := vmids[dockerVMID]; !ok {
		missing = append(missing, fmt.Sprintf("docker.vmid %d not on node", dockerVMID))
	}
	if len(missing) > 0 {
		return fmt.Errorf("%s", strings.Join(missing, "; "))
	}
	return nil
}

// ListGuests returns LXC and QEMU guests, sorted by ID. TYPE qemu is mapped to vm.
func ListGuests(ctx context.Context, runner Runner) ([]Guest, error) {
	node, err := discoverNode(ctx, runner)
	if err != nil {
		return nil, err
	}

	out := make([]Guest, 0)
	seen := make(map[int]struct{})
	kinds := []struct {
		pathKind string
		cliType  string
	}{
		{"lxc", "lxc"},
		{"qemu", "vm"},
	}
	for _, k := range kinds {
		path := fmt.Sprintf("/nodes/%s/%s", node, k.pathKind)
		raw, err := pveshGet(ctx, runner, path)
		if err != nil {
			return nil, err
		}
		var items []guestListItem
		if err := json.Unmarshal(raw, &items); err != nil {
			return nil, fmt.Errorf("parse %s json: %w", path, err)
		}
		for _, item := range items {
			if item.VMID <= 0 {
				return nil, fmt.Errorf("guest missing vmid on %s", path)
			}
			if item.Status == "" {
				return nil, fmt.Errorf("guest %d missing status", item.VMID)
			}
			if item.Status != "running" && item.Status != "stopped" {
				return nil, fmt.Errorf("guest %d: unknown status %q", item.VMID, item.Status)
			}
			if _, dup := seen[item.VMID]; dup {
				return nil, fmt.Errorf("duplicate vmid %d", item.VMID)
			}
			seen[item.VMID] = struct{}{}
			name := item.Name
			if strings.ContainsAny(name, "\t\n") {
				return nil, fmt.Errorf("guest %d: name contains tab or newline", item.VMID)
			}
			if name == "" {
				name = "-"
			}
			out = append(out, Guest{
				ID:    item.VMID,
				Type:  k.cliType,
				Name:  name,
				Power: item.Status,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// CheckSSH runs a trivial pvesh call to prove connectivity and auth.
func CheckSSH(ctx context.Context, runner Runner) error {
	_, err := pveshGet(ctx, runner, "/nodes")
	return err
}
