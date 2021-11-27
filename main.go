package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/webvillain/vikashbank13/handlers"
	"github.com/webvillain/vikashbank13/model"
)

func main() {
	os.Remove("./bank.db")
	fmt.Println("Application Is Starting On Port : 9001 on Your Local Machine")
	model.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/mybank", handlers.ListUsers).Methods("GET")
	r.HandleFunc("/mybank/{id}", handlers.UserById).Methods("GET")
	r.HandleFunc("/mybank/{firstname}/{lastname}/{email}", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/mybank/{id}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/mybank/{id}/{email}", handlers.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9001", r))
}
