package data

import (
	"go_songs/models"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Collection struct {
	Cl *mgo.Collection
}

func (col Collection) FindAll() ([]models.Song, error) {
	var data []models.Song
	err := col.Cl.Find(nil).All(&data)
	return data, err
}

func (col Collection) FindById(id string) (models.Song, error) {
	var data models.Song
	err := col.Cl.FindId(id).One(&data)
	return data, err
}

func (col Collection) CountId(id string) int {
	n, _ := col.Cl.FindId(id).Count()
	return n
}

func (col Collection) MaxId() int {
	var data models.Song
	_ = col.Cl.Find(nil).Sort("-_id").One(&data)
	n, _ := strconv.Atoi(data.Id)
	return n
}

func (col Collection) Insert(data models.Song) error {
	err := col.Cl.Insert(&data)
	return err
}

func (col Collection) Delete(id string) error {
	err := col.Cl.Remove(bson.M{"_id": id})
	return err
}

func (col Collection) Update(id string, k bson.M) error {
	err := col.Cl.UpdateId(id, bson.M{"$set": k})
	return err
}
