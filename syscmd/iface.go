// Description: This file contains the interface related functions.
package syscmd

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"

	log "github.com/golang/glog"
)

// InternetConnectedInterface returns the network interface used to connect to the internet.
func InternetConnectedInterface() (*net.Interface, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	iface, err := net.InterfaceByName(localAddr.Zone)
	if err != nil {
		iface, err = interfaceByLocalAddress(localAddr.IP)
		if err != nil {
			return nil, err
		}
	}

	return iface, nil
}

func interfaceByLocalAddress(ip net.IP) (*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ipNet *net.IPNet
			switch v := addr.(type) {
			case *net.IPNet:
				ipNet = v
			case *net.IPAddr:
				ipNet = &net.IPNet{
					IP:   v.IP,
					Mask: net.CIDRMask(128, 128),
				}
			}

			if ipNet.Contains(ip) {
				return &iface, nil
			}
		}
	}
	return nil, errors.New("no interface found for the IP")
}

// Get the WiFi driver name used by a specific interface
func (r *Runner) GetDriver(interfaceName string) (string, error) {
	driverPath := fmt.Sprintf("/sys/class/net/%s/device/driver", interfaceName)
	driverRealPath, err := filepath.EvalSymlinks(driverPath)
	if err != nil {
		return "", err
	}
	return filepath.Base(driverRealPath), nil
}

// BringUp brings up a certain network interface.
func (r *Runner) BringUp(intfName string) error {
	if _, err := r.ExecCommand(true, "ip link set dev", intfName, "up"); err != nil {
		return err
	}
	log.Infof("Interface %v is UP.", intfName)
	return nil
}

// TurnDown turns down a certain network interface.
func (r *Runner) TurnDown(intfName string) error {
	if _, err := r.ExecCommand(true, "ip link set dev", intfName, "down"); err != nil {
		return err
	}
	log.Infof("Interface %v is DOWN.", intfName)
	return nil
}

// SetCountry sets the country code for a certain network interface.
func (r *Runner) SetCountry(intfName string, countryCode string) error {
	if _, err := r.ExecCommand(true, "iw", intfName, "set", "country", countryCode); err != nil {
		return err
	}
	log.Infof("Country code %v is set for interface %v.", countryCode, intfName)
	return nil
}

// SetTxPower sets the tx power for a certain network interface.
func (r *Runner) SetTxPower(intfName string, txPower int) error {
	if _, err := r.ExecCommand(true, "iw", intfName, "txpower", "fixed", fmt.Sprintf("%d", txPower)); err != nil {
		return err
	}
	log.Infof("Tx power %v is set for interface %v.", txPower, intfName)
	return nil
}

// SetUnmanaged sets a certain network interface as unmanaged.
func (r *Runner) SetUnmanaged(intfName string) error {
	if _, err := r.ExecCommand(true, "nmcli dev set", intfName, "managed", "no"); err != nil {
		return err
	}
	log.Infof("Interface %v is set as unmanaged.", intfName)
	return nil
}

// SetManaged sets a certain network interface as managed.
func (r *Runner) SetManaged(intfName string) error {
	if _, err := r.ExecCommand(true, "nmcli dev set", intfName, "managed", "yes"); err != nil {
		return err
	}
	log.Infof("Interface %v is set as managed.", intfName)
	return nil
}

// SetMAC sets the MAC address for a certain network interface.
func (r *Runner) SetMAC(intfName string, macAddr string) error {
	if _, err := r.ExecCommand(true, "ip link set dev", intfName, "address", macAddr); err != nil {
		return err
	}
	log.Infof("MAC address %v is set for interface %v.", macAddr, intfName)
	return nil
}

// Reset a certain network interface.
func (r *Runner) Reset(intfName string) error {
	if _, err := r.ExecCommand(true, "ip", "link", "set", "dev", intfName, "down"); err != nil {
		return err
	}
	if _, err := r.ExecCommand(true, "ip link set dev", intfName, "up"); err != nil {
		return err
	}
	log.Infof("Interface %v is reset.", intfName)
	return nil
}
