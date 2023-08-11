package main

import (
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Define target IP
	dstIP := net.IPv4(192, 168, 0, 1)

	// Open a raw network device for sending packets
	handle, err := pcap.OpenLive("eth0", 1600, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create ICMP Echo Request layer
	icmp := &layers.ICMPv4{
		TypeCode: layers.CreateICMPv4TypeCode(layers.ICMPv4TypeEchoRequest, 0),
	}

	// Create IP layer
	ip := &layers.IPv4{
		SrcIP:    net.IPv4(192, 168, 0, 2), // Replace with source IP
		DstIP:    dstIP,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolICMPv4,
	}

	// Create Ethernet layer
	eth := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}, // Replace with source MAC
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // Broadcast or replace with destination MAC
		EthernetType: layers.EthernetTypeIPv4,
	}

	// Serialize the packet
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buf, opts, eth, ip, icmp)
	if err != nil {
		log.Fatal(err)
	}

	// Send the packet
	err = handle.WritePacketData(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Packet sent!")
}

