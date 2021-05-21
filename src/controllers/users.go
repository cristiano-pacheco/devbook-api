package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
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
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Update"))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Delete"))
}
