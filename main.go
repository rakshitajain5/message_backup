package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"message_backup/controllers"
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
	var i int
	test(&i)
	fmt.Println(i)
	fmt.Fprintf(w, "Welcome to go, %q", html.EscapeString(r.URL.Path))
}

func test(i *int) {
	*i = 10
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
