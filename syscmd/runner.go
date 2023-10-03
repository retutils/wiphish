// Package syscmd provides methods executing external commands.
package syscmd

import (
	"os"
	"os/exec"

	log "github.com/golang/glog"
)

// Runner contains methods executing external commands.
type Runner struct {
	// ExecCommand runs the given command with arguments.
	// It returns the content of stdout or stderr, and error if command failed.
	ExecCommand func(wait bool, cmd string, args ...string) (string, error)
}

func execute(wait bool, cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)
	var output []byte
	var err error

	if wait {
		output, err = command.CombinedOutput()
	} else {
		command.Stdout = os.Stdout
		err = command.Start()
	}

	outputString := string(output)
	if err != nil {
		log.Errorf("Command (%v %v) failed. Error: %v.\nOutput:\n%v", cmd, args, err, outputString)
	} else {
		log.V(2).Infof("Command (%v %v) succeeded.\nOutput:\n%v", cmd, args, outputString)
	}

	return outputString, err
}

// SetIPForwading enables IP forwarding.
func (r *Runner) SetIPForwading() error {
	if _, err := r.ExecCommand(true, "sysctl", "-w", "net.ipv4.ip_forward=1"); err != nil {
		return err
	}
	log.Info("IP forwarding is enabled.")
	return nil
}

// OffIPForwading disables IP forwarding.
func (r *Runner) OffIPForwading() error {
	if _, err := r.ExecCommand(true, "sysctl", "-w", "net.ipv4.ip_forward=0"); err != nil {
		return err
	}
	log.Info("IP forwarding is disabled.")
	return nil
}
func New() *Runner {
	return &Runner{
		ExecCommand: execute}
}
