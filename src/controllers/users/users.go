package users

import (
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"ping": "pong"})
}

func GetOne(w http.ResponseWriter, r *http.Request) {

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

	responses.Success(w, http.StatusCreated, user)

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
