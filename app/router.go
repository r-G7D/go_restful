package app

import (
	"r-G7D/go_restful/handler"

	"github.com/julienschmidt/httprouter"
)

func DefRoute(userHandler handler.UserHandler) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", userHandler.FindAll)
	router.GET("/api/users/:id", userHandler.FindById)
	router.POST("/api/users", userHandler.Create)
	router.PUT("/api/users/:id", userHandler.Update)
	router.DELETE("/api/users/:id", userHandler.Delete)

	return router
}
