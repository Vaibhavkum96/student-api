package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vaibhavkum96/student-api-go/internal/config"
)

func main() {
	fmt.Println("Welcome to Students Api Project in Go!")

	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome To Students Api!"))
	})
	// setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}
	fmt.Printf("Server Started! %s", cfg)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed To Start Server!")
	}

}
