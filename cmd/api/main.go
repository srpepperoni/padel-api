package main

import (
	"os"

	"k8s.io/klog/v2"

	"fake.com/padel-api/config"
	"fake.com/padel-api/internal/server"
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

	cfg, _ := config.GetConfig(os.Args[1])

	r := server.NewRouter(cfg)
	server.Run(cfg, r)
}
