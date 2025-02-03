package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/harsh-jagtap-josh/RozgarLink/internal/api"
	_ "github.com/lib/pq"
)

func main() {
	mux := mux.NewRouter()
	// db, err := repository.ConnectDB()
	// if err != nil {
	// 	fmt.Println("Database error Handler")
	// 	return
	// }
	// -- Routes --
	routes.NewRouter(mux)
	fmt.Println("Server up and running on Port : 8080 :)")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
