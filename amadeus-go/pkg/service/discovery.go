package service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulSd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"amadeus-go/config/env"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Register , registers a service in consul discovery service
func Register(consulAddress string, serviceAddress string, httpPort string, grpcPort string) sd.Registrar {
	var registrar sd.Registrar
	// Logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	
	rand.Seed(time.Now().UTC().UnixNano())
	
	// Service discovery
	var client consulSd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulSd.NewClient(consulClient)
	}
	
	check := api.AgentServiceCheck{
		HTTP:     "http://" + serviceAddress + ":" + httpPort + "/health",
		Interval: "60s",
		Timeout:  "15s",
		Notes:    "Basic health checks",
	}
	
	port, _ := strconv.Atoi(grpcPort)
	asr := api.AgentServiceRegistration{
		ID:      env.GetString("service.name"),
		Name:    env.GetString("service.name"),
		Address: serviceAddress,
		Port:    port,
		Tags:    env.GetStringSlice("service.discovery_tags"),
		Check:   &check,
	}
	
	registrar = consulSd.NewRegistrar(client, &asr, logger)
	return registrar
}

// GetServerAddress , returns the address of gRPC service that already has been registered in consul
// Using the service name and one of the service tags
func GetServerAddress(consulAddr, serviceName, tag string) (string, error) {
	
	config := api.DefaultConfig()
	
	config.Address = consulAddr          // Assign address
	client, err := api.NewClient(config) // Creating consul client
	if err != nil {
		return "", err
	}
	
	// Getting service info from consul agent
	services, _, err := client.Health().Service(serviceName, tag, false, nil)
	if err != nil || len(services) == 0 {
		return "", err
	}
	
	// Extracting address and port from info
	serviceAddress := services[0].Service.Address
	servicePort := services[0].Service.Port
	
	// Return address in address:port format
	return fmt.Sprintf("%s:%d", serviceAddress, servicePort), nil
}
