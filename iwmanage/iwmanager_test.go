package iwmanage

import "testing"

const (
	testIntfName = "wlan1"
	testTxPower  = 10
)

func TestGetWiFiDriver(t *testing.T) {
	cmd := Command()
	if _, err := cmd.GetWiFiDriver(testIntfName); err != nil {
		t.Errorf("Getting WiFi driver failed. Error: %v.", err)
	}
}

func TestCheckHostapdSupport(t *testing.T) {
	cmd := Command()
	if _, err := cmd.CheckHostapdSupport(testIntfName); err != nil {
		t.Errorf("Checking hostapd support failed. Error: %v.", err)
	}
}
