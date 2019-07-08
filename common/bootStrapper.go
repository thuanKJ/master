package common

import (
	"github.com/globalsign/mgo"
)

//StartUp
func StartUp() (*mgo.Collection, *mgo.Collection) {
	colSong, colAccount := initmongoDB()
	return colSong, colAccount
}
