// package cfg provide configuration for the application
package cfg

import (
	"os"
	"path/filepath"

	"github.com/wiphish/syscmd"
)

// Config is the configuration variables for the application
type Config struct {
	TempDir string //Temporarily directory to store temporary config files
	// HostapdConfFile is the file path to the hostapd configuration file.
	HostapdConfFile string
	// HostapdLogFile specifies the path to the log file for hostapd.
	HostapdLogFile string
	// ExtInterface  is the name of shared Internet of the external interface to use.
	ExtInterface string
	// WiFiInterface is the name of the wireless interface to be used for network configuration.
	WiFiInterface string
	// WiFiSSID is the name of the WiFi network to connect to.
	WiFiSSID string
	// WiFiBSSID is the BSSID of the WiFi network.
	WiFiBSSID string
	// TempDir is the directory to store temporary config files
	Network string
	//DhcpPool is the ip range of dhcp pool
	DhcpPool string
	// Evilginx is the IP address of the evilginx server.
	Evilginx string
	// Channel is the WiFi channel to use.
	Channel uint
	// Gateway is the IP address of the gateway.
	Gateway string
	// IptablesRulesFile is the file to save itpables rule.
	IptablesRulesFile string
	// DhcpConfigFile is the file to save dhcp config.
	DhcpConfigFile string
	// CMD is the command to execute.
	CMD *syscmd.Runner
}

func New() (cfg *Config, err error) {
	cfg = &Config{}
	if cfg.TempDir, err = os.MkdirTemp("", "running"); err != nil {
		return nil, err
	}
	cfg.HostapdConfFile = filepath.Join(cfg.TempDir, "hostapd.conf")
	cfg.HostapdLogFile = filepath.Join(cfg.TempDir, "hostapd.log")
	cfg.IptablesRulesFile = filepath.Join(cfg.TempDir, "iptables.rules")
	cfg.DhcpConfigFile = filepath.Join(cfg.TempDir, "dhcpd.conf")
	cfg.CMD = syscmd.Command()
	return cfg, nil
}

func (cfg *Config) Clean() error {
	return os.RemoveAll(cfg.TempDir)
}

func (cfg *Config) SetDhcpPool(pool string) {
	cfg.DhcpPool = pool
}

func (cfg *Config) SetNetwork(network string) {
	cfg.Network = network
}

func (cfg *Config) SetGateway(gateway string) {
	cfg.Gateway = gateway
}

func (cfg *Config) SetChannel(channel uint, err error) {
	cfg.Channel = channel
}

func (cfg *Config) SetWiFiInterface(wifiInterface string) {
	cfg.WiFiInterface = wifiInterface
}

func (cfg *Config) SetWiFiSSID(wifiSSID string) {
	cfg.WiFiSSID = wifiSSID
}

func (cfg *Config) SetWiFiBSSID(wifiBSSID string) {
	cfg.WiFiBSSID = wifiBSSID
}

func (cfg *Config) SetExtInterface(extInterface string) {
	cfg.ExtInterface = extInterface
}
