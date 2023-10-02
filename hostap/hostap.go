// Description: This file contains the functions to configure , start and stop the hostapd service.
package hostap

import (
	"fmt"
	"strings"

	log "github.com/golang/glog"
	"github.com/wiphish/syscmd"
)

type Runner struct {
	cmd *syscmd.Runner
}

// Check if the driver of a given interface supports hostapd
func (r *Runner) CheckHostapdSupport(interfaceName string) (bool, error) {
	output, err := r.cmd.ExecCommand(true, "iw", interfaceName, "info")
	if err != nil {
		return false, fmt.Errorf("failed to retrieve information about %s: %v", interfaceName, err)
	}
	return strings.Contains(output, "* AP\n"), err
}

// StartHostapd starts the hostapd service with the specified configuration file.
// The configFile parameter is the path to the hostapd configuration file
func (r *Runner) StartHostapd(configFile string) error {
	if _, err := r.cmd.ExecCommand(false, "hostapd", configFile); err != nil {
		return err
	}
	log.Infof("Started a hostapd with config file: %v.", configFile)
	return nil
}

// StopHostapd stops the hostapd process link to the given WLAN interface.
func (r *Runner) StopHostapd() error {
	if _, err := r.cmd.ExecCommand(true, "killall", "-q", "hostapd"); err != nil {
		return err
	}
	log.Info("Stopped all hostapd processes.")
	return nil
}

func New() *Runner {
	return &Runner{
		cmd: syscmd.Command(),
	}
}
