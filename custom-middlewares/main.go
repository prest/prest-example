package main

import (
	"fmt"
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
	n := middlewares.GetApp()

	// Register custom middleware
	n.UseFunc(customMiddleware)

	// Get pPREST router
	r := router.Get()

	// Register custom routes
	r.HandleFunc("/ping", pong).Methods("GET")

	// Call pREST cmd
	cmd.Execute()
}

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pong!"))
	if err != nil {
		fmt.Println(err)
	}
}

func customMiddleware(
	w http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc) {
	_, err := fmt.Fprintln(w, "Calling custom middleware")
	if err != nil {
		fmt.Println(err)
		return
	}
	next(w, r)
}
