package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func PublicationGetAll(w http.ResponseWriter, r *http.Request) {
// 	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repository := repositories.NewPublicationRepository(db)
// 	publications, err := repository.Search(nameOrNick)

// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	responses.JSON(w, http.StatusOK, publications)
// }

func PublicationGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

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

	repository := repositories.NewPublicationRepository(db)
	publication, err := repository.Get(id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

func PublicationCreate(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	publication.AuthorID = userID

	if err = publication.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewPublicationRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

// func PublicationUpdate(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	ID, err := strconv.ParseUint(params["id"], 10, 64)

// 	if err != nil {
// 		responses.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	tokenPublicationID, err := authentication.ExtractPublicationId(r)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	if ID != tokenPublicationID {
// 		responses.Error(w, http.StatusForbidden, errors.New("is not possible to update an publication that is not you"))
// 		return
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	var publication models.Publication
// 	if err = json.Unmarshal(body, &publication); err != nil {
// 		responses.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	if err = publication.Prepare("update"); err != nil {
// 		responses.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	defer db.Close()

// 	repository := repositories.NewPublicationRepository(db)

// 	err = repository.Update(ID, publication)
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	responses.JSON(w, http.StatusNoContent, nil)
// }

// func PublicationDelete(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	ID, err := strconv.ParseUint(params["id"], 10, 64)

// 	if err != nil {
// 		responses.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	tokenPublicationID, err := authentication.ExtractPublicationId(r)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	if ID != tokenPublicationID {
// 		responses.Error(w, http.StatusForbidden, errors.New("is not possible to delete an publication that is not you"))
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	defer db.Close()

// 	repository := repositories.NewPublicationRepository(db)

// 	err = repository.Delete(ID)
// 	if err != nil {
// 		responses.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	responses.JSON(w, http.StatusNoContent, nil)
// }
