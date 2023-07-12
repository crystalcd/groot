package main

import (
	"log"
	"net/http"
	"time"

	"github.com/crystal/groot/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           ":8089",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
