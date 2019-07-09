package controllers

import (
	"go_songs/common"
	"go_songs/models"
	"encoding/json"
	"io"
	"net/http"
)

var colAccount common.ConnectionData

// Get session
func GetAccountSession(typeData common.ConnectionData) {
	colAccount = typeData
	return
}

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Acc models.Account
	err := json.NewDecoder(r.Body).Decode(&Acc)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(`Please fill Username and Password`)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := colAccount.CheckByName(Acc.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Account does not exist")
	} else {
		account, err := colAccount.FindByName(Acc.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if Acc.Username == account.Username && Acc.Password == account.Password {
			TokenString, err := common.GenerateJWT(Acc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				json.NewEncoder(w).Encode(TokenString)
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode("Worng Password")
		}
	}
}

func AddNewAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Acc models.Account
	err := json.NewDecoder(r.Body).Decode(&Acc)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please fill Username and Password")
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := colAccount.CheckByName(Acc.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username exists")
	} else {
		err = colAccount.AddNewAccount(Acc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Sign up successfully")
	}
}
