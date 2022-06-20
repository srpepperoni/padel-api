package main

import (
	"fmt"
	"log"
	"net/http"

	router "fake.com/padel-api/internal/server"
)

func main() {
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router.NewRouter()))
}
