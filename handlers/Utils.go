package handlers

import (
	"log"
	"net/http"
)

// checkErrHTTP check handler errors and sends HTTP response 500
func checkErrHTTP(err error, writer http.ResponseWriter, message string) {
	if err != nil {
		log.Println(message, err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
