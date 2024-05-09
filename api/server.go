package api

import (
	"github.com/marcos-travasso/simple-api/core/repository"
	"github.com/marcos-travasso/simple-api/core/repository/in_memory"
	core "github.com/marcos-travasso/simple-api/core/service"
	"log"
	"net/http"
	"os"
)

var service *core.Service

func StartServer() {
	repo := in_memory.NewRepository()
	CreateService(repo)

	http.HandleFunc("POST /reset", resetHandler)
	http.HandleFunc("GET /balance", getBalanceHandler)
	http.HandleFunc("POST /event", postEventHandler)

	port := ":8080"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(port, nil))
}

func GetService() *core.Service {
	return service
}

func CreateService(repo repository.Repository) {
	service = core.NewService(repo)
}
