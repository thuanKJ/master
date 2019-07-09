package common

import (
	"fmt"
	"go_songs/models"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type ConnectionData interface {
	FindAll() ([]models.Song, error)
	FindById(id string) (models.Song, error)
	CountId(id string) int
	MaxId() int
	Insert(data models.Song) error
	Delete(id string) error
	Update(id string, k bson.M) error
	FindByName(name string) (models.Account, error)
	AddNewAccount(Account models.Account) error
	CheckByName(name string) (bool, error)
}

// Dao just dao
type Dao struct {
	Server   string
	Database string
}

const (
	collectionSong    = "song"
	collectionAccount = "account"
)

var server Dao
var db *mgo.Database

func initmongoDB() (*mgo.Collection, *mgo.Collection) {
	server.Server = "db:27017"
	server.Database = "music"
	colSong, colAccount := server.Conect()
	return colSong, colAccount
}

func (m *Dao) Conect() (*mgo.Collection, *mgo.Collection) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		fmt.Println("Error mongo connection")
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	return db.C(collectionSong), db.C(collectionAccount)
}
