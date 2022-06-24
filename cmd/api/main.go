package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fake.com/padel-api/config"
	router "fake.com/padel-api/internal/server"
	"fake.com/padel-api/pkg/utils"
)

// @title Go padel-api
// @version 1.0
// @description Golang Rest API for padel tournaments
// @contact.name Jaime Yera
// @contact.url https://github.com/srpepperoni
// @contact.email jaimeyera@gmail.com
func main() {
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)

	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	fmt.Printf("Server at %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, router.NewRouter(cfg)))
}
