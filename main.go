package main

import (
	"log"
	"net/http"

	"github.com/DouglasMarq/go-deployer/handlers"
)

func main() {
	http.HandleFunc("/deploy", handlers.DeployHandler)

	log.Println("Server started on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
