package main

import (
	"github.com/marcos-travasso/simple-api/api"
	"github.com/marcos-travasso/simple-api/core/repository/in_memory"
	"github.com/marcos-travasso/simple-api/core/service"
)

func main() {
	repo := in_memory.NewRepository()
	s := service.NewService(repo)

	api.StartServer(s)
}
