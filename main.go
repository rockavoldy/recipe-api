package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rockavoldy/recipe-api/category"
	"github.com/rockavoldy/recipe-api/unit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var HTTP_PORT = "8000"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dsn := "host=db user=recipe password=recipe dbname=recipes_api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when open connection to DB: %v\n", err)
	}

	category.SetDB(db)
	unit.SetDB(db)
	r.Mount("/category", category.Router())
	r.Mount("/unit", unit.Router())

	log.Printf("Listening on port :%s\n", HTTP_PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", HTTP_PORT), r); err != nil {
		log.Fatalf("HTTP server error: %v\n", err)
	}
}
