// cmd/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"double-dose/discovery/internal/discovery"
)

func main() {
	sd, err := discovery.NewServiceDiscovery()
	if err != nil {
		log.Fatalf("Failed to create service discovery: %v", err)
	}
	defer sd.DeregisterService("order-management")

	serviceID := "order-management"
	servicePort := 8080
	err = sd.RegisterService("order-management", serviceID, "localhost", servicePort)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("Shutting down...")
}
