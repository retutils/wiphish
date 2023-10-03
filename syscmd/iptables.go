// Package iptables provides the functions to configure , start and stop the iptables service.
package syscmd

import (
	"os"

	log "github.com/golang/glog"
)

// Temporaly file handeler to store iptables rules
var iptfile *os.File

// Temoraly dirrectory to store iptables rules file
var tempDir *os.DirEntry

// IptablesFlush the iptables rules.
func (r *Runner) IptablesFlush() error {
	if _, err := r.ExecCommand(true, "iptables", "-F"); err != nil {
		return err
	}
	log.Info("Runner rules are flushed.")
	return nil
}

// IptablesAcceptAll accepts all packets.
func (r *Runner) IptablesAcceptAll() error {
	if _, err := r.ExecCommand(true, "iptables", "-P", "INPUT", "ACCEPT"); err != nil {
		return err
	}
	if _, err := r.ExecCommand(true, "iptables", "-P", "OUTPUT", "ACCEPT"); err != nil {
		return err
	}
	if _, err := r.ExecCommand(true, "iptables", "-P", "FORWARD", "ACCEPT"); err != nil {
		return err
	}
	log.Info("Runner rules are set to accept all packets.")
	return nil
}

// IptablesMasqueradeAll masquerades all packets.
// The intfName parameter is the name of the interface to masquerade.
// It is usually the interface connected to the Internet.
func (r *Runner) IptablesMasqueradeAll(intfName string) error {
	if _, err := r.ExecCommand(true, "iptables", "-t", "nat", "-A", "POSTROUTING", "-o", intfName, "-j", "MASQUERADE"); err != nil {
		return err
	}
	log.Infof("Runner rules are set to masquerade all packets from interface %v.", intfName)
	return nil
}

// RunnerRestore restore iptables rules from a file.
func (r *Runner) Restore(file string) error {
	if _, err := r.ExecCommand(true, "iptables-restore", file); err != nil {
		return err
	}
	log.Infof("Runner rules are restored from file %v.", file)
	return nil
}

// Save saves iptables rules to a file.
func (r *Runner) Save(file string) error {
	if _, err := r.ExecCommand(true, "iptables-save", ">", file); err != nil {
		return err
	}
	log.Infof("Runner rules are saved to file %v.", file)
	return nil
}
