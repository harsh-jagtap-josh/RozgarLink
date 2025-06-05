package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"

	db "github.com/harsh-jagtap-josh/RozgarLink"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {

	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		logger.Errorw(ctx, "failed to load .env variables", zap.Error(err))
		return
	}

	sqlDB, err := db.InitDB(ctx)
	if err != nil {
		logger.Errorw(ctx, "failed to establish a database connection", zap.Error(err))
		return
	}

	services := app.NewServices(sqlDB)
	router := app.NewRouter(services)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	})
	handler := cors.Default().Handler(router)

	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: c.Handler(handler),
	}

	logger.Infow(ctx, "Server up and running", zap.String("port", string(port)))
	err = srv.ListenAndServe()
	if err != nil {
		logger.Errorw(ctx, "failed to start server", zap.Error(err), zap.String("port", string(port)))
		return
	}
}
