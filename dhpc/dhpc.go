package dhcpcd

import (
	log "github.com/golang/glog"
	"github.com/wiphish/syscmd"
)

type Runner struct {
	cmd *syscmd.Runner
}

// StartDHCPCD starts the dhcpcd service.
// /sbin/dhcpd -f -cf /etc/dhcp/dhcpd.conf -user dhcpd -group dhcpd --no-pid eth0
func (r *Runner) StartDHCPCD() error {
	if _, err := r.cmd.ExecCommand(true, "systemctl", "start", "dhcpcd"); err != nil {
		return err
	}
	log.Info("Dhcpcd is started.")
	return nil
}

// StopDHCPCD stops the dhcpcd service.
func (r *Runner) StopDHCPCD() error {
	if _, err := r.cmd.ExecCommand(true, "systemctl", "stop", "dhcpcd"); err != nil {
		return err
	}
	log.Info("Dhcpcd is stopped.")
	return nil
}

// StopDnsmasq stops the dhcpcd service.
func (r *Runner) StopDnsmnasq() error {
	if _, err := r.cmd.ExecCommand(true, "systemctl", "stop", "dnsmasq"); err != nil {
		return err
	}
	log.Info("Dhcpcd is stopped.")
	return nil
}

// StartDnsmasq starts the dhcpcd service.
// /sbin/dhcpd -f -cf /etc/dhcp/dhcpd.conf -user dhcpd -group dhcpd --no-pid eth0
func (r *Runner) StartDnsmasq() error {
	if _, err := r.cmd.ExecCommand(true, "systemctl", "start", "dnsmasq"); err != nil {
		return err
	}
	log.Info("Dhcpcd is started.")
	return nil
}

// StopAvahi stops the avahi-daemon service.
func (r *Runner) StopAvahi() error {
	if _, err := r.ExecCommand(true, "systemctl", "stop", "avahi-daemon"); err != nil {
		return err
	}
	log.Info("Avahi-daemon is stopped.")
	return nil
}

// StartAvahi starts the avahi-daemon service.
func (r *Runner) StartAvahi() error {
	if _, err := r.ExecCommand(true, "systemctl", "start", "avahi-daemon"); err != nil {
		return err
	}
	log.Info("Avahi-daemon is started.")
	return nil
}
e