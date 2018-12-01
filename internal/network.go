package internal

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// GetIP returns the ip address for the given interface prefixes
// prefers a v4 address and will return the first one it finds
func GetIP(prefixes []string) string {
	var myip string

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		// Make sure the interface has a prefix we care about
		hasPrefix := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(iface.Name, prefix) && iface.Name != "wlan0:0" {
				hasPrefix = true
			}
		}

		// Grab the addresses on the interface and make sure
		// its not a loopback or anything we don't want
		if hasPrefix {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				ip, _, _ := net.ParseCIDR(addr.String())
				if !ip.IsLoopback() {
					if ip.To4() == nil {
						continue
					}

					if myip == "" && ip.String() != "172.23.86.75" {
						myip = ip.String()
					}
				}
			}
		}
	}

	return myip
}

// CopyNetworkConfig copies our source wlan0 config and puts it in
// /etc/network/interfaces.d/wlan0 we need one that has the wpa_supplicant
// commented out so that it doesn't wreck hostapd, then one with it for when
// we reboot and try to connect (that change still to come)
func CopyNetworkConfig(src string) {
	in, err := os.Open(src)
	if err != nil {
		log.Fatal("Could not open source network config", err.Error())
	}
	defer in.Close()

	out, err := os.Create("/etc/network/interfaces.d/wlan0")
	if err != nil {
		log.Fatal("Could not create /etc/network/interfaces.d/wlan0", err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatal("Could not copy contents to wlan0", err.Error())
	}
}
