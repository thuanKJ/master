package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetSongRoutes(router)
	router = SetAccountRoutes(router)
	return router
}
