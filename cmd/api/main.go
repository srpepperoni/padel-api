package main

import (
	"k8s.io/klog/v2"
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
	klog.InitFlags(nil)
	defer klog.Flush()
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)

	if err != nil {
		klog.Fatal("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)

	if err != nil {
		klog.Fatalf("ParseConfig: %v", err)
	}

	klog.Infof("Server at %s", cfg.Server.Port)
	klog.Fatal(http.ListenAndServe(cfg.Server.Port, router.NewRouter(cfg)))
}
