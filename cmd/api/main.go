package main

import (
	"fmt"
	"log"
	"net/http"

	router "fake.com/padel-api/internal/server"
)

func main() {
	fmt.Println("Server at 8000")
	//TODO: change the port to an env variable
	log.Fatal(http.ListenAndServe(":8000", router.NewRouter()))
}
