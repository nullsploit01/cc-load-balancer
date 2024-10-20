package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	ServerName          string `mapstructure:"server_name"`
	URL                 string `mapstructure:"url"`
	HealthCheckURL      string `mapstructure:"health_check_url"`
	HealthCheckInterval int    `mapstructure:"health_check_interval"`
	IsHealthy           bool
}

var serverStateLock sync.Mutex

func main() {
	config, err := LoadConfig()

	if err != nil {
		panic(err)
	}

	for i := range config.Servers {
		go startHealthCheck(&config.Servers[i])
	}

	for {
		server, err := chooseHealthyServer(config.Servers)
		if err == nil {
			fmt.Printf("Forwarding request to: %s\n", server.URL)
		} else {
			fmt.Println("No healthy servers available!")
		}
		time.Sleep(5 * time.Second) // Simulate requests coming in
	}
}

func startHealthCheck(server *Server) {
	ticker := time.NewTicker(time.Duration(server.HealthCheckInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		response, err := http.Get(server.HealthCheckURL)
		serverStateLock.Lock()
		if err != nil || response.StatusCode != http.StatusOK {
			server.IsHealthy = false
			fmt.Printf("Server %s is DOWN\n", server.ServerName)
		} else {
			server.IsHealthy = true
			fmt.Printf("Server %s is UP\n", server.ServerName)
		}
		serverStateLock.Unlock()
	}
}

var lastUsedServer = -1

func chooseHealthyServer(servers []Server) (Server, error) {
	serverStateLock.Lock()
	defer serverStateLock.Unlock()
	for i := 0; i < len(servers); i++ {
		lastUsedServer = (lastUsedServer + 1) % len(servers)
		if servers[lastUsedServer].IsHealthy {
			return servers[lastUsedServer], nil
		}
	}

	return Server{}, fmt.Errorf("No healthy servers available")
}
