package discovery

import (
	"log"

	"github.com/hashicorp/consul/api"
)

type ServiceDiscovery struct {
	client *api.Client
}

func NewServiceDiscovery() (*ServiceDiscovery, error) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ServiceDiscovery{client: client}, nil
}

func (sd *ServiceDiscovery) RegisterService(serviceName, serviceID, serviceAddress string, servicePort int) error {
	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		ID:      serviceID,
		Address: serviceAddress,
		Port:    servicePort,
	}

	err := sd.client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	log.Printf("Registered service: %s", serviceName)
	return nil
}

func (sd *ServiceDiscovery) DeregisterService(serviceID string) error {
	err := sd.client.Agent().ServiceDeregister(serviceID)
	if err != nil {
		return err
	}

	log.Printf("Deregistered service: %s", serviceID)
	return nil
}
