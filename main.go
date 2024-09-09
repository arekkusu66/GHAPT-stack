package main

import (
	"app/templates/components"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	staticPath  = "/static"
	portEnv     = "PORT"
	defaultPort = "8080"
)

func serveStaticFiles() {
	fileServer := http.FileServer(http.Dir("." + staticPath))
	http.Handle(staticPath + "/", http.StripPrefix(staticPath + "/", fileServer))
}

func getPort() string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}
	return ":" + port
}

func main() {
	components.Routes()

	serveStaticFiles()

	port := getPort()
	fmt.Printf("Listening on %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
