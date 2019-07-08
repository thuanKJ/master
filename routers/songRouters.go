package routers

import (
	"go_songs/controllers"

	"github.com/gorilla/mux"
)

func SetSongRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/songs", controllers.GetAllData).Methods("GET")
	router.HandleFunc("/song/{id}", controllers.GetOne).Methods("GET")
	router.HandleFunc("/song", controllers.CreateOne).Methods("POST")
	router.HandleFunc("/song/{id}", controllers.DeleteOne).Methods("DELETE")
	router.HandleFunc("/song/{id}", controllers.UpdateOne).Methods("PATCH")
	return router
}
