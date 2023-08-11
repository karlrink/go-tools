package main

import (
	"fmt"
	"net"
	"time"
)

func ping(ip string) {
	// Set up a deadline for the connection attempt
	conn, err := net.DialTimeout("ip4:icmp", ip, 1*time.Second)
	if err != nil {
		fmt.Printf("Host %s is down: %s\n", ip, err)
		return
	}
	defer conn.Close()

	// Set a read deadline (timeout)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	// Send an ICMP Echo Request message (type 8)
	conn.Write([]byte{8, 0, 0, 0, 0, 0, 0, 0})

	// Read the ICMP Echo Reply message (type 0)
	buf := make([]byte, 32)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("Host %s is down: %s\n", ip, err)
		return
	}

	// Check if the reply message type is 0 (Echo Reply)
	if buf[0] == 0 {
		fmt.Printf("Host %s is up\n", ip)
	} else {
		fmt.Printf("Received unexpected message from %s\n", ip)
	}
}

func main() {
	// Loop through the desired IP range to scan multiple hosts
	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("192.168.0.%d", i)
		ping(ip)
	}
}

