package syscmd

import (
	log "github.com/golang/glog"
)

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
