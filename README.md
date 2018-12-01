# PI Network Helper

A simple tool that will check if your pi (really any linux device that has systemd and the prereqs) and if it is not, will start hostapd and turn it into an access point you can then connect to and configure.

If it is on the network, can be given a service name to start.

This was made to ease taking my [DrinkMachine](https://github.com/neophenix/drinkmachine) to friend's houses and not have to take a keyboard and monitor to connect and setup wifi, and all that entails

## Prerequisites

1. Raspbian or some other OS with systemd
1. hostapd
1. dnsmasq

## Usage

Get your Pi on your network as normal.  This assumes the configuration will live at /etc/network/interfaces.d/wlan0.  One catch, we are going to use wlan0 as our AP interface, and wlan0:0 as our normal network interface.  We do this because, I believe it was easier to get the configuration right, and I think using wlan0 for both stepped on the toes of normal networking and hostapd, sadly its been awhile since I first wrote this and didn't write down the reasons.

Next, configure hostapd (/etc/hostapd/hostapd.conf) really just setting your passphrase.

Then, configure dnsmasq (/etc/dnsmasq.conf) setting the network range you will be using.

For the above 3 files, samples can be found in this repo.

Now, we are going to want an "AP" version of our wlan0 config, you can use the provided wlan0.ap.sample.  Place this file somewhere that the service will be able to read it.  When the service determines we are not connected to a network, it will put this file in place of the real wlan0 and restart networking.  Ideally, this file is the same as your normal wlan0 file except that the link to wpa-supplicant.conf is commented out for wlan0:0 so it dosn't try to connect to anything.

Build the tool: ```go build -o SOMENAME main.go```

Using the provided systemd definition sample, add the service to systemd and enable it so it starts at boot.  Change the options to match your usage:

```
-interfaces : A comma seperated list of interface prefixes to look for an IP (default: wlan,eth)
-services : A comma seperated list of services to start if we have an IP, no default
-apconf : Location of the AP version of wlan0 (the one with wpa_supplicant likely commented out) (default: wlan0.ap)
```

That should be it, on boot the service will check if it is on the network and start the service(s) or go into AP mode.
