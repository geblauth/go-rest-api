package routes

import (
	"log"
	"net/http"

	"github.com/geblauth/GoRestApi/controllers"
	"github.com/geblauth/GoRestApi/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
