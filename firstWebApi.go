package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {

	// ... / api

	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		// w.Header().Set("Content-Type","application/json")
		fmt.Fprintf(w, string(output))
	})

	// ... / api / users

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "John", LastName: "Connor", Age: 25},
			User{ID: 2, FirstName: "Thomas", LastName: "William", Age: 32},
			User{ID: 3, FirstName: "Samantha", LastName: "O'Brien", Age: 28},
			User{ID: 4, FirstName: "Charles", LastName: "Miller", Age: 22},
			User{ID: 5, FirstName: "Patricia", LastName: "Tremblay", Age: 19},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// ... / api / me

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{3, "Yusuf", "ALTUN", 28}
		message := user
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
