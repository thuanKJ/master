package controllers

import (
	"encoding/json"
	"go_songs/common"
	"go_songs/models"
	"io"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var colSong common.ConnectionData

//Get session
func GetSongSession(typeData common.ConnectionData) {
	colSong = typeData
	return
}

// GetAllData - Get all item from mongoDB
func GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := colSong.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// GetOne - Get a item from mongoDB
func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	data, err := colSong.FindById(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// CreateOne - Create a item from mongoDB
func CreateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please type down data")
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	n := colSong.MaxId()
	song.Id = strconv.Itoa(n + 1)
	err = colSong.Insert(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(song)
}

// DeleteOne - Delete a item from mongoDB
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := colSong.Delete(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(bson.M{"remove": "OK"})
}

// UpdateOne
func UpdateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if colSong.CountId(params["id"]) != 0 {
		var k bson.M
		err := json.NewDecoder(r.Body).Decode(&k)
		switch {
		case err == io.EOF:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Please type down data")
			return
		case err != nil:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = colSong.Update(params["id"], k)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := colSong.FindById(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Doesn't exist Id")
	}
}
