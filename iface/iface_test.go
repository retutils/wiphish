package iface

import (
	"testing"
)

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
