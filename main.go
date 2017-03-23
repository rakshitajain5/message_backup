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
	http.Server{}.SetKeepAlivesEnabled(true)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go, %q", html.EscapeString(r.URL.Path))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//func msgBackup(w http.ResponseWriter, r *http.Request) {
//	//fmt.Fprintf(w, "Hello World")
//	decoder := json.NewDecoder(r.Body)
//	deviceKey := r.Header.Get("X-Device-Key")
//	if deviceKey=="" {
//		http.Error(w, "X-Device-Key missing", http.StatusBadRequest)
//		return
//	}
//	fmt.Println(string(decoder))
//}