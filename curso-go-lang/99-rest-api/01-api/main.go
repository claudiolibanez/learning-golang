package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("Server is running in 8080")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {

}
