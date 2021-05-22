package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(nameOrNick)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
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
	user, err := repository.Get(uint64(id))

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
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

	if err = user.Prepare("create"); err != nil {
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
	user.Password = ""
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if ID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to update an user that is not you"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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

	err = repository.Update(ID, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if ID != tokenUserID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to delete an user that is not you"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	err = repository.Delete(ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UserFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	followerID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to follow you"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	err = repository.Follow(userID, followerID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

func UserUnFollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	followerID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("is not possible to unfollow you"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	err = repository.UnFollow(userID, followerID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
