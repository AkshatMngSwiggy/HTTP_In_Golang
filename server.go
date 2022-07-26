package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	fmt.Println("Server started listening...")
	http.HandleFunc("/", GetHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	u.Firstname = "Akshat"
	u.Lastname = "Mangal"

	user_acc := User{}

	switch r.Method {
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&user_acc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "User Firstname: %s\n", user_acc.Firstname)
		fmt.Fprintf(w, "User Lastname: %s\n", user_acc.Lastname)
		fmt.Fprintf(w, "Post request Completed\n")
	case "GET":
		err := json.NewEncoder(w).Encode(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Get request Completed\n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
