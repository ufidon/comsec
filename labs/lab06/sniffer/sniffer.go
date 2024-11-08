package main

// sudo apt-get install libpcap-dev
// go get -u github.com/google/gopacket
// go get -u github.com/google/gopacket/pcap
// # Set CGO_ENABLED to 0 to build a static binary
// CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o sniffer sniffer.go
// sudo setcap 'cap_net_raw,cap_net_admin,cap_sys_nice=eip' ./sniffer_static
// getcap sniffer
// ./sniffer --interface enp0s3 --timeout 10000000

import (
	"flag"
	"fmt"
	"log"

	// "os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Define command-line flags for network interface and timeout
	device := flag.String("interface", "eth0", "Network interface to capture packets from (e.g., eth0, wlan0)")
	timeout := flag.Int("timeout", 30, "Capture timeout in seconds")
	flag.Parse()

	// Open device for capturing
	snapshotLen := int32(1024) // Size of the packet capture
	promiscuous := false       // Set to true if you want all packets on the network
	timeoutDuration := time.Duration(*timeout) * time.Second
	handle, err := pcap.OpenLive(*device, snapshotLen, promiscuous, timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use gopacket to capture packets in a loop
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Starting packet capture...")

	// Iterate through packets
	for packet := range packetSource.Packets() {
		processPacket(packet)
	}
}

func processPacket(packet gopacket.Packet) {
	// Print packet timestamp and capture length
	fmt.Printf("\nTime: %s\n", packet.Metadata().Timestamp)
	fmt.Printf("Packet Length: %d bytes\n", packet.Metadata().CaptureLength)

	// Print network layer info
	networkLayer := packet.NetworkLayer()
	if networkLayer != nil {
		fmt.Printf("Source IP: %s\n", networkLayer.NetworkFlow().Src())
		fmt.Printf("Destination IP: %s\n", networkLayer.NetworkFlow().Dst())
	}

	// Print transport layer info
	transportLayer := packet.TransportLayer()
	if transportLayer != nil {
		fmt.Printf("Protocol: %s\n", transportLayer.LayerType())
	}

	// Dump entire packet data if desired
	fmt.Println("Full Packet:", packet.Dump())
}
