package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luanrem/GoGiota-API/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOGIOTA_DATABASE_USER"),
		os.Getenv("GOGIOTA_DATABASE_PASSWORD"),
		os.Getenv("GOGIOTA_DATABASE_HOST"),
		os.Getenv("GOGIOTA_DATABASE_PORT"),
		os.Getenv("GOGIOTA_DATABASE_NAME"),
	))
	if err != nil {
		panic(err) // o erro do pgxpool.New não pega se o db está no ar, só se está malformada
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.Api{
		Router: chi.NewMux(),
	}

	api.BindRoutes()

	fmt.Println("Server is running on :3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}
