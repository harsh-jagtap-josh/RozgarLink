package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/api"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// ctx := context.Background()

	godotenv.Load()
	sqlDB, err := repository.InitDB()
	if err != nil {
		fmt.Println("error connecting to database" + err.Error())
		return
	}

	services := app.NewServices(sqlDB)
	router := api.NewRouter(services)

	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router,
	}
	fmt.Println("Server up and running on PORT", port)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Errorf("error while starting http server, %s", err.Error())
		return
	}
}
