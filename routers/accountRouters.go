package routers

import (
	"go_songs/controllers"

	"github.com/gorilla/mux"
)

func SetAccountRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/signin", controllers.Signin).Methods("GET")
	router.HandleFunc("/signup", controllers.AddNewAccount).Methods("POST")
	return router
}
