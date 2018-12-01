package internal

import (
	"log"
	"os/exec"
)

// StartAPServices starts dnsmasq and hostapd to put the Pi in Access Point mode.
func StartAPServices() {
	log.Println("Restarting networking")
	cmd := exec.Command("systemctl", "restart", "networking")
	cmd.Run()

	log.Println("Starting dnsmasq")
	cmd = exec.Command("systemctl", "start", "dnsmasq")
	cmd.Run()

	log.Println("Starting hostapd")
	cmd = exec.Command("systemctl", "start", "hostapd")
	cmd.Run()
}

// StartServices takes a list of services provided by the user on the command line -services flag and
// systemctl starts them
func StartServices(services []string) {
	for _, service := range services {
		log.Println("Starting", service)
		cmd := exec.Command("systemctl", "start", service)
		cmd.Run()
	}
}
