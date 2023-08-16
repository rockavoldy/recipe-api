package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rockavoldy/recipe-api/category"
	"github.com/rockavoldy/recipe-api/material"
	"github.com/rockavoldy/recipe-api/recipe"
	"github.com/rockavoldy/recipe-api/recipematerial"
	"github.com/rockavoldy/recipe-api/unit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var HTTP_PORT = "8000"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dbconfig := dbConfig{}
	dbconfig.loadFromEnv()
	db, err := gorm.Open(postgres.Open(dbconfig.ConnStr()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when open connection to DB: %v\n", err)
	}

	category.SetDB(db)
	unit.SetDB(db)
	material.SetDB(db)
	recipe.SetDB(db)
	recipematerial.SetDB(db)
	r.Mount("/category", category.Router())
	r.Mount("/unit", unit.Router())
	r.Mount("/material", material.Router())
	r.Mount("/recipe", recipe.Router())

	loadEnvStr("RECIPE_HTTP_PORT", &HTTP_PORT)
	log.Printf("Listening on port :%s\n", HTTP_PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", HTTP_PORT), r); err != nil {
		log.Fatalf("HTTP server error: %v\n", err)
	}
}
