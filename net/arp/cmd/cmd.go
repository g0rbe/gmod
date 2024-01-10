/*
 * Example CLI for arp package.
 * Try to find every device on the interface's subnet.
 *
 * Build: go build .
 *
 * Run: sudo ./cmd
 */
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/g0rbe/gmod/net/arp"
	"github.com/g0rbe/gmod/net/iface"
	"github.com/g0rbe/gmod/net/ip"
)

func main() {

	if os.Geteuid() != 0 {
		fmt.Fprintf(os.Stderr, "RUN AS ROOT!\n")
		os.Exit(1)
	}

	devs, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get inteerfaces: %s\n", err)
		os.Exit(1)
	}

	for i := range devs {

		fmt.Printf("Interface: %s\n", devs[i].Name)

		if devs[i].Name == "lo" {
			fmt.Printf("\tSkipping loopback...\n")
			continue
		}

		nets, err := iface.GetIPNets4(&devs[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "\tFailed to get %s IPv4 networks: %s\n", devs[i].Name, err)
			continue
		}

		for v := range nets {

			fmt.Printf("\tSubnet: %s\n", nets[v].String())

			first, last, usable, err := ip.GetList(*nets[v])
			if err != nil {
				fmt.Fprintf(os.Stderr, "\t\tFailed to list ip addresses: %s\n", err)
				continue
			}

			fmt.Printf("\t\tFirst address: %s\n", first)
			fmt.Printf("\t\tLast address: %s\n", last)

			for ip := range usable {

				fmt.Printf("\t\tSearching for %s -> ", ip)

				entry, err := arp.Get(ip, 2*time.Second)
				if err != nil {
					if strings.Contains(err.Error(), "i/o timeout") {
						fmt.Printf("No device found\n")
					} else {
						fmt.Printf("Error for %s: %s\n", ip, err)
					}
					continue
				}

				fmt.Printf("%s\n", entry.Addr)

				if err := arp.DelARPCache(entry); err != nil {
					fmt.Printf("\t\tFailed to delete ARP cache for %s: %s\n", ip, err)
				}
			}
		}
	}
}
