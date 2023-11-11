package handlers

import (
	"encoding/json"
	"github.com/office-sports/odds-api/data"
	"github.com/office-sports/odds-api/models"
	"log"
	"net/http"
)

func GenerateToken(writer http.ResponseWriter, request *http.Request) {
	var tr models.TokenRequest
	err := json.NewDecoder(request.Body).Decode(&tr)

	checkErrHTTP(err, writer, "Unable to parse input data")

	user, err := data.GetUserByLogin(tr.Login)
	if err != nil {
		log.Println("Invalid user")
		http.Error(writer, "Invalid user", http.StatusBadRequest)
		return
	}

	err = user.CheckPassword(tr.Password)
	if err != nil {
		log.Println("Invalid password")
		http.Error(writer, "Invalid password", http.StatusForbidden)
		return
	}

	tokenString, err := auth.GenerateJWT(user.Login)
	if err != nil {
		log.Println("Error generating token")
		http.Error(writer, "Error generating token", http.StatusInternalServerError)
		return
	}

	t := new(models.Token)
	t.Token = tokenString

	json.NewEncoder(writer).Encode(t)
}
