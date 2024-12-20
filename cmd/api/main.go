package main

import(
	"fmt"
	"net/http"
	// Import 'chi' as framework
	"github.com/go-chi/chi"
	// Importing from internal/handlers folder
	"github.com/min-verse/receipt-processor/internal/handlers"
	// Importing logrus to use as a logging service
	log "github.com/sirupsen/logrus"
)

func main(){
	// Setting up Logging service
	log.SetReportCaller(true)

	// Setting up the router via chi
	var r *chi.Mux = chi.NewRouter()

	// Defined in internals/handlers
	// This will set up our router
	// which adds the endpoint definitions
	handlers.Handler(r)

	// Start up message
	fmt.Println("Starting Receipt Processor GO API Server...")

	err := http.ListenAndServe(":8080", r)
	// Logs any errors when server fails to start
	if err != nil{
		log.Error(err)
	}
}