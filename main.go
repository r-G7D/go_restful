package main

import (
	"net/http"
	"r-G7D/go_restful/app"
	"r-G7D/go_restful/handler"
	"r-G7D/go_restful/helper"
	"r-G7D/go_restful/repository"
	"r-G7D/go_restful/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.DefDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userHandler := handler.NewUserHandler(userService)

	router := app.DefRoute(userHandler)

	server := http.Server{
		Addr:    "localhost:8008",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
