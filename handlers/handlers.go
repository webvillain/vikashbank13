package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/webvillain/vikashbank13/model"
)

var db *sql.DB
var Store = model.NewDatabse(db)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List Of All Users")
	w.Header().Set("Content-Type", "application/json")
	users, err := Store.ListUser()
	if err != nil {
		log.Fatal(err)
	}
	res := json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusFound)
	fmt.Fprintln(w, res)

}
func UserById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get User By Id")
	opts := mux.Vars(r)
	id := opts["id"]
	newid, _ := strconv.ParseInt(id, 0, 0)
	user, err := Store.SingleUser(newid)
	if err != nil {
		log.Fatal(err)
	}
	res := json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusFound)
	fmt.Fprintln(w, res)

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create New User")
	opts := mux.Vars(r)
	fistname := opts["firstname"]
	lastname := opts["lastname"]
	email := opts["email"]
	err := Store.CreateUser(fistname, lastname, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Is Created Successfully.")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete a User")
	opts := mux.Vars(r)
	id := opts["id"]
	newid, _ := strconv.ParseInt(id, 0, 0)
	err := Store.DeleteUser(newid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "User Is Deleted Successfully.")

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update a User")
	opts := mux.Vars(r)
	id := opts["id"]
	firstname := opts["firstname"]
	lastname := opts["lastname"]
	email := opts["email"]
	newid, _ := strconv.ParseInt(id, 0, 0)
	if err := Store.UpdateUser(newid, firstname, lastname, email); err == nil {
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("User Is Updated Successfully.")
	// to do

}
