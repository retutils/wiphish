package syscmd

import (
	log "github.com/golang/glog"
	"github.com/wiphish/syscmd"
)

// Runner is a dhcpcd runner.
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
