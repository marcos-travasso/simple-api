package api

import (
	"fmt"
	core "github.com/marcos-travasso/simple-api/core/service"
	"log"
	"net/http"
	"os"
)

var service *core.Service

func StartServer(s *core.Service) {
	InjectService(s)

	http.HandleFunc("POST /reset", resetHandler)
	http.HandleFunc("GET /balance", getBalanceHandler)
	http.HandleFunc("POST /event", postEventHandler)

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	log.Println(fmt.Sprintf("Listening at %s", port))
	log.Fatal(http.ListenAndServe(port, nil))
}

func GetService() *core.Service {
	return service
}

func InjectService(s *core.Service) {
	service = s
}
