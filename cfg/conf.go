// package cfg provide configuration for the application
package cfg

import "github.com/wiphish/iptables"

// Config is the configuration variables for the application
type Config struct {
	TempDir         string             //Temporarily directory to store temporary config files
	Itpables        *iptables.Iptables //Iptables cmds
	HostapdConfFile string
	HostapdLogFile  string
	ExtInterface    string
	WiFiInterface   string
	WiFiSSID        string
	WiFiBSSID       string
}
