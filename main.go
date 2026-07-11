package main

import (
	"log"
	"net/http"

	"http-server-projeto-korp/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	
	// Configura todas as rotas usando o pacote interno
	routes.Setup(mux)

	log.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
