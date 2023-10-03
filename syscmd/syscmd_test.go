package syscmd

import (
	"testing"
)

const (
	testIntfName = "wlan1"
)

func TestCheckHostapdSupport(t *testing.T) {
	r := Command()
	if _, err := r.CheckHostapdSuppor(testIntfName); err != nil {
		t.Errorf("Checking hostapd support failed. Error: %v.", err)
	}
}

// You would ideally mock net.Dial and other dependencies, but for this example, we're testing the real function.
// The following test assumes you have an active internet connection.
func TestInternetConnectedInterface(t *testing.T) {
	iface, err := InternetConnectedInterface()
	if err != nil {
		t.Fatalf("Failed to determine internet-connected interface: %v", err)
	}

	if iface == nil {
		t.Fatal("Expected an interface, but got nil")
	}
	t.Log(iface.Name)
}
