package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func pingIP(ip string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("ping", "-c", "3", ip)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Host %s is down\n", ip)
	} else {
		fmt.Printf("Host %s is up\n", ip)
	}
}

func main() {
	var wg sync.WaitGroup

	// Loop through the desired IP range to scan multiple hosts
	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("192.168.0.%d", i)
		wg.Add(1)
		go pingIP(ip, &wg)
	}

	wg.Wait()
	fmt.Println("Scan finished.")
}

