package main

import (
	"fmt"
	"log"
	"net/http"

	router "fake.com/padel-api/internal/server"
)

// @title Go padel-api
// @version 1.0
// @description Golang Rest API for padel tournaments
// @contact.name Jaime Yera
// @contact.url https://github.com/srpepperoni
// @contact.email jaimeyera@gmail.com
func main() {
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router.NewRouter()))
}
