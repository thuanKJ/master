package main

import (
	"go_songs/common"
	"go_songs/controllers"
	"go_songs/data"
	"go_songs/routers"

	"github.com/codegangsta/negroni"
)

func main() {
	//Create a session connect to MongoDB
	clSong, clAccount := common.StartUp()
	song := data.Collection{clSong}
	account := data.Collection{clAccount}

	//Get connection in controller
	controllers.GetSongSession(song)
	controllers.GetAccountSession(account)

	//Create router
	router := routers.InitRouters()

	//Add middleware
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(common.Validate))
	n.UseHandler(router)
	n.Run(":8000")
}
