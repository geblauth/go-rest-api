package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/geblauth/GoRestApi/models"
	"github.com/gorilla/mux"
)

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
func RetornaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {

		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
