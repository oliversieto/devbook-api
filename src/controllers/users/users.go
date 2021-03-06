package users

import (
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"strconv"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("user"))

	dbConnection, err := database.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)
	users, err := repository.GetAll(search)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.Success(w, http.StatusOK, users)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConnection, err := database.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)
	user, err := repository.GetOne(ID)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.Success(w, http.StatusOK, user)

}

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(true); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConnection, err := database.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	user.Phrase = ""

	responses.Success(w, http.StatusCreated, user)

}

func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(false); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConnection, err := database.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)

	if err = repository.Update(user, ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.Success(w, http.StatusNoContent, nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	dbConnection, err := database.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)

	if err = repository.Delete(ID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.Success(w, http.StatusNoContent, nil)
}
