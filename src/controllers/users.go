package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Get All"))
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Get"))
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", userID)))
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Update"))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Delete"))
}
