package hostap

import (
	"testing"

	"github.com/wiphish/syscmd"
)

const (
	testIntfName = "wlan1"
)

func TestCheckHostapdSupport(t *testing.T) {
	r := &Runner{cmd: syscmd.Command()}
	if _, err := r.CheckHostapdSupport(testIntfName); err != nil {
		t.Errorf("Checking hostapd support failed. Error: %v.", err)
	}
}
