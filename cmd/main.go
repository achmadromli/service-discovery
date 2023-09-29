// cmd/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"double-dose/discovery/internal/discovery"
	"double-dose/discovery/internal/service"
)

func main() {
	// Buat service discovery
	sd, err := discovery.NewServiceDiscovery()
	if err != nil {
		log.Fatalf("Failed to create service discovery: %v", err)
	}
	defer sd.DeregisterService("my-service")

	// Daftarkan layanan ke service discovery
	serviceID := "my-service"
	servicePort := 8080
	err = sd.RegisterService("my-service", serviceID, "localhost", servicePort)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	// Inisialisasi dan jalankan layanan
	myService := service.NewMyService(servicePort)
	go myService.Start()

	// Tunggu sinyal SIGINT atau SIGTERM untuk menghentikan layanan
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("Shutting down...")
}
