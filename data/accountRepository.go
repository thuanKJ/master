package data

import (
	"fmt"
	"go_songs/models"

	"github.com/globalsign/mgo/bson"
)

func (col Collection) FindByName(name string) (models.Account, error) {
	var account models.Account
	err := col.Cl.Find(bson.M{"username": name}).One(&account)
	return account, err
}

func (col Collection) AddNewAccount(account models.Account) error {
	obj_id := bson.NewObjectId()
	account.Id = obj_id
	err := col.Cl.Insert(account)
	return err
}

func (col Collection) CheckByName(name string) (bool, error) {
	v, err := col.Cl.Find(bson.M{"username": name}).Count()
	fmt.Println("name", v)
	if v == 0 {
		return false, err
	} else {
		return true, err
	}
}
