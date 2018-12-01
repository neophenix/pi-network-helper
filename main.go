package main

import (
	"flag"
	"github.com/neophenix/pi-helpers/internal"
	"log"
	"strings"
	"time"
)

func main() {
	var prefixes, services, apconf string
	flag.StringVar(&prefixes, "interfaces", "wlan,eth", "A comma seperated list of interface prefixes to look for an IP")
	flag.StringVar(&services, "services", "", "A comma seperated list of services to start if we have an IP")
	flag.StringVar(&apconf, "apconf", "wlan0.ap", "Location of the AP version of wlan0 (the one with wpa_supplicant likely commented out)")
	flag.Parse()
	myip := internal.GetIP(strings.Split(prefixes, ","))

	internal.SetupLCD()
	defer internal.CloseLCD()

	if myip == "" {
		log.Println("Entering AP Mode")
		internal.WriteString("Entering AP Mode", 0, 0)
		internal.CopyNetworkConfig(apconf)
		internal.StartAPServices()
		// give us time to see the info
		time.Sleep(5 * time.Second)
	} else {
		log.Printf("Connected to network with address %v", myip)
		internal.WriteString("IP Address", 0, -1)
		internal.WriteString(myip, 1, -1)
		// give us time to see the info
		time.Sleep(5 * time.Second)
		if services != "" {
			internal.StartServices(strings.Split(services, ","))
		}
	}

	internal.CloseLCD()
}
