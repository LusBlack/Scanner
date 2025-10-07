package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanPort(port int, wg *sync.WaitGroup) {
	IP := "scanme.nmap.org"
	defer wg.Done()
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)

	if err != nil {
		return
	}
	fmt.Printf("[+] Connection established..PORT %v ", port)
	connection.Close()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go ScanPort(i, &wg)
	}
	wg.Wait()
}
