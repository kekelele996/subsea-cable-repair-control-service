package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kekelele996/subsea-cable-repair-control-service/internal/api"
	"github.com/kekelele996/subsea-cable-repair-control-service/internal/service"
)

func main() {
	addr := os.Getenv("CABLE_REPAIR_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	system := service.NewSystem(time.Now)
	server := api.NewServer(system)
	log.Printf("cable repair control listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatal(err)
	}
}
