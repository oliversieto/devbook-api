package auth

import (
	"devbook-api/src/authentication"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"devbook-api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request) {
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
	savedUser, err := repository.GetByEmail(user.Email)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePhrases(savedUser.Phrase, user.Phrase); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.TokenGenerator(savedUser.ID)

	responses.Success(w, http.StatusOK, token)
}
