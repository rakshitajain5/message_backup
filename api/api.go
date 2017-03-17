package api

import (
	"github.com/gorilla/mux"
	"message_backup/controllers"
	"net/http"
	"fmt"
	"html"
)

func Handlers() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", Index).Methods("POST")
	router.HandleFunc("/", Hello).Methods("GET")
	router.HandleFunc("/jcm/messages/backup", controllers.MsgBackup).Methods("POST")
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go, %q", html.EscapeString(r.URL.Path))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}