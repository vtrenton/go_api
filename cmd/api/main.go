package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/vtrenton/go_api/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	// welcome messages
	fmt.Println("Starting Go API service...")
	fmt.Println(`
  _________    ___   ___  ____
 / ___/ __ \  / _ | / _ \/  _/
/ (_ / /_/ / / __ |/ ___// /  
\___/\____/ /_/ |_/_/  /___/ `)

	// Start the http server and catch any errors into the err variable
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}

}
