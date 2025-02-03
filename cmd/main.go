package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/api"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
)

func main() {
	// ctx := context.Background()

	sqlDB, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("error connecting to database" + err.Error())
		return
	}

	services := app.NewServices(sqlDB)

	router := api.NewRouter(services)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}
	fmt.Println("Server up and running on port 8080")
	log.Fatal(srv.ListenAndServe())
}
