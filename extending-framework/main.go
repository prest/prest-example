package main

import (
	"net/http"

	"github.com/prest/adapters/postgres"
	"github.com/prest/cmd"
	"github.com/prest/config"
	"github.com/prest/config/router"
	"github.com/prest/middlewares"
)

func main() {
	config.Load()

	// Load Postgres Adapter
	postgres.Load()

	// Get pREST app
	middlewares.GetApp()

	// Get pPREST router
	r := router.Get()

	// Register custom routes
	r.HandleFunc("/ping", pong).Methods("GET")

	// Call pREST cmd
	cmd.Execute()
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!"))
}
