package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Read in the config
	cfg := ReadConfig()

	// Register the route functions
	http.HandleFunc("/", Router)

	// Output and start serving
	log.Println(fmt.Sprintf("Serving on %s:%s", cfg.Network.Interface, cfg.Network.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Network.Interface, cfg.Network.Port), nil))
}
