package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
)

func must(desc string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", desc, err)
	}
}

func main() {
	// Initialize the Bluetooth adapter.
	d, err := linux.NewDevice()
	must("new device", err)
	//defer d.Stop()

	// Set a timeout for scanning (e.g., 10 seconds).
	timeout := 100 * time.Second

	// Start scanning for BLE devices.
	fmt.Println("Scanning for devices...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Discover devices during the scan.
	err = d.Scan(ctx, false, func(a ble.Advertisement) {
		val := ""
		if len(a.Services()) > 0 {
			val = fmt.Sprint(a.Services()[0])

			if val == "1819" {
				fmt.Printf("Device name:%s\t SSID:%s\n", a.LocalName(), val)

			}
		}

	})

	must("scan", err)

	fmt.Println("Scan completed.")
}
