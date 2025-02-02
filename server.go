package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/createUser", createUser)

	http.ListenAndServe(":8080", nil)
}

func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to the GO server!")
}

// Encodes (marshals) the users slice into a JSON object and writes it to the response writer.
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// Decodes (unmarshals) it into the newUser variable.
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	// add id to the new user
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
