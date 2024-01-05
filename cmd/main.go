package main

import (
	"log"
	config "ssh-proxy-app/config"
	client "ssh-proxy-app/internal/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	client.Run(*conf)
}
