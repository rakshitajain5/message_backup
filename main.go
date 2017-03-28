package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"message_backup/controllers"
	"github.com/gorilla/mux"
)

type messages struct {

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("POST")
	router.HandleFunc("/hello", Hello).Methods("GET")
	router.HandleFunc("/jcm/messages/backup", controllers.MsgBackup).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go, %q", html.EscapeString(r.URL.Path))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

