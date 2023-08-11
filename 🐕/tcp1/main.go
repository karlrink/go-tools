package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Millisecond*500)

	if err == nil {
		fmt.Printf("Found open port: %d at IP: %s\n", port, ip)
		conn.Close()
	}
}

func main() {
	var wg sync.WaitGroup

	// IP address to scan
	ip := "192.168.0.1"

	// Range of ports to scan
	for port := 20; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(ip, port, &wg)
	}

	wg.Wait()
	fmt.Println("Scan finished.")
}

