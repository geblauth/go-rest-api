package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func main() {

	handlerequest()

}
