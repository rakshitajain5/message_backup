package main

import _ "net/http/pprof"

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"message_backup/controllers"
	"log"
)

type messages struct {

}

func main() {
	fmt.Println("Server starting")
	r := http.DefaultServeMux
	r.HandleFunc("/jcm/messages/backup", controllers.MsgBackup)
	r.HandleFunc("/", Index)
	//log.Fatal(http.ListenAndServe(":8080", api.Handlers()))
	log.Fatal(http.ListenAndServe(":8080", r))
}



func Handlers() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("POST")
	router.HandleFunc("/hello", Hello).Methods("GET")
	router.HandleFunc("/jcm/messages/backup", controllers.MsgBackup).Methods("POST")
	return router
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