package routes

import (
	"log"
	"net/http"

	"github.com/geblauth/GoRestApi/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
