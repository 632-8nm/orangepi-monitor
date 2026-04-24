package main

import "os"

func main() {
	server := NewServer()
	addr := os.Getenv("MONITOR_LISTEN_ADDR")
	if addr == "" {
		// Bind all interfaces by default so LAN access works out of the box.
		addr = "0.0.0.0:8080"
	}
	server.Start(addr)
}
