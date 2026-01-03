package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"go-auth-api/db"
	"go-auth-api/handlers"
	"go-auth-api/middleware"
)

func main() {
	godotenv.Load()
	db.Connect()
	r := chi.NewRouter()

	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)

	r.With(middleware.AuthMiddleware).Get("/profile", handlers.Profile)

	log.Printf("Server Running On Port:", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
