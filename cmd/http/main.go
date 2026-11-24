package main

import (
	"go-wallet/internal/shared/src/interfaces/http"
	"log"
)

func main() {
	r := http.BuildApp()
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start HTTP server: ", err)
	}
}
