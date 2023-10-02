// Package iptables provides the functions to configure , start and stop the iptables service.
package iptables

import (
	"os"

	log "github.com/golang/glog"
	"github.com/wiphish/cfg"
	"github.com/wiphish/syscmd"
)

type Iptables struct {
	cmd           *syscmd.Runner
	ExtInterface  string
	WifiInterface string
}

// Temporaly file handeler to store iptables rules
var iptfile *os.File

// Temoraly dirrectory to store iptables rules file
var tempDir *os.DirEntry

// IptablesFlush flushes the iptables rules.
func (r *Iptables) Flush() error {
	if _, err := r.cmd.ExecCommand(true, "iptables", "-F"); err != nil {
		return err
	}
	log.Info("Iptables rules are flushed.")
	return nil
}

// AcceptAll accepts all packets.
func (r *Iptables) AcceptAll() error {
	if _, err := r.cmd.ExecCommand(true, "iptables", "-P", "INPUT", "ACCEPT"); err != nil {
		return err
	}
	if _, err := r.cmd.ExecCommand(true, "iptables", "-P", "OUTPUT", "ACCEPT"); err != nil {
		return err
	}
	if _, err := r.cmd.ExecCommand(true, "iptables", "-P", "FORWARD", "ACCEPT"); err != nil {
		return err
	}
	log.Info("Iptables rules are set to accept all packets.")
	return nil
}

// MasqueradeAll masquerades all packets.
// The intfName parameter is the name of the interface to masquerade.
// It is usually the interface connected to the Internet.
func (r *Iptables) MasqueradeAll(intfName string) error {
	if _, err := r.cmd.ExecCommand(true, "iptables", "-t", "nat", "-A", "POSTROUTING", "-o", intfName, "-j", "MASQUERADE"); err != nil {
		return err
	}
	log.Infof("Iptables rules are set to masquerade all packets from interface %v.", intfName)
	return nil
}

// IptablesRestore restore iptables rules from a file.
func (r *Iptables) Restore(file string) error {
	if _, err := r.cmd.ExecCommand(true, "iptables-restore", file); err != nil {
		return err
	}
	log.Infof("Iptables rules are restored from file %v.", file)
	return nil
}

// Save saves iptables rules to a file.
func (r *Iptables) Save(file string) error {
	if _, err := r.cmd.ExecCommand(true, "iptables-save", ">", file); err != nil {
		return err
	}
	log.Infof("Iptables rules are saved to file %v.", file)
	return nil
}

// Init initializes iptables rules.
func New(c *cfg.Config) (err error) {
	return nil
}
