package users

import (
	"devbook-api/repositories"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	dbConnection, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	repository := repositories.NewUserRepository(dbConnection)
	ID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]uint64{"id": ID})
}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
