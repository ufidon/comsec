package main

import (
	"flag"
	"fmt"
	"log"

	// "os"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Struct to count packets from each IP to detect port scans
type packetCount struct {
	Count    int
	LastSeen time.Time
}

var (
	alertMutex sync.Mutex
	packetMap  = make(map[string]*packetCount)
)

func main() {
	// Define command-line flags for network interface and timeout
	device := flag.String("interface", "eth0", "Network interface to capture packets from (e.g., eth0, wlan0)")
	timeout := flag.Int("timeout", 30, "Capture timeout in seconds")
	flag.Parse()

	// Open device for capturing
	snapshotLen := int32(1024)
	promiscuous := false
	timeoutDuration := time.Duration(*timeout) * time.Second
	handle, err := pcap.OpenLive(*device, snapshotLen, promiscuous, timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use gopacket to capture packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("Starting packet capture...")

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
		srcIP := networkLayer.NetworkFlow().Src().String()
		dstIP := networkLayer.NetworkFlow().Dst().String()
		fmt.Printf("Source IP: %s\n", srcIP)
		fmt.Printf("Destination IP: %s\n", dstIP)

		// Check for potential port scanning
		checkPortScan(srcIP)
	}

	// Print transport layer info
	transportLayer := packet.TransportLayer()
	if transportLayer != nil {
		fmt.Printf("Protocol: %s\n", transportLayer.LayerType())

		if tcpLayer, ok := transportLayer.(*layers.TCP); ok {
			// Detect SYN flood by checking SYN flags without ACK
			if tcpLayer.SYN && !tcpLayer.ACK {
				fmt.Println("Alert: SYN packet detected without ACK (potential SYN flood attempt)")
			}

			fmt.Printf("Source Port: %d\n", tcpLayer.SrcPort)
			fmt.Printf("Destination Port: %d\n", tcpLayer.DstPort)
		}

		if udpLayer, ok := transportLayer.(*layers.UDP); ok {
			fmt.Printf("UDP Source Port: %d\n", udpLayer.SrcPort)
			fmt.Printf("UDP Destination Port: %d\n", udpLayer.DstPort)
		}
	}

	// Dump packet data if necessary
	fmt.Println("Full Packet:", packet.Dump())
}

func checkPortScan(srcIP string) {
	// Track packet count from each source IP
	alertMutex.Lock()
	defer alertMutex.Unlock()

	if packetMap[srcIP] == nil {
		packetMap[srcIP] = &packetCount{Count: 1, LastSeen: time.Now()}
	} else {
		packetMap[srcIP].Count++
		packetMap[srcIP].LastSeen = time.Now()

		if packetMap[srcIP].Count > 10 {
			fmt.Printf("Alert: Potential port scan detected from IP %s (more than 10 packets)\n", srcIP)
			// Reset count after alert
			packetMap[srcIP].Count = 0
		}
	}

	// Clean up old entries to prevent memory leak
	for ip, count := range packetMap {
		if time.Since(count.LastSeen) > time.Minute {
			delete(packetMap, ip)
		}
	}
}
