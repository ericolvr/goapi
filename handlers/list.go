package handlers

import (
	"encoding/json"
	"gingo/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao aobter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
