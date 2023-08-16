package main

import (
	"fmt"
	"os"
	"strconv"
)

type dbConfig struct {
	Host   string
	Port   uint
	DbName string
	DbUser string
	DbPass string
}

func (dbc dbConfig) ConnStr() string {
	// dsn := "host=db.recipe-api.orb.local user=recipe password=recipe dbname=recipes_api port=5432 sslmode=disable"
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", dbc.Host, dbc.Port, dbc.DbName, dbc.DbUser, dbc.DbPass)
}

func (dbc *dbConfig) loadFromEnv() {
	loadEnvStr("RECIPE_DB_HOST", &dbc.Host)
	loadEnvUint("RECIPE_DB_PORT", &dbc.Port)
	loadEnvStr("RECIPE_DB_NAME", &dbc.DbName)
	loadEnvStr("RECIPE_DB_USER", &dbc.DbUser)
	loadEnvStr("RECIPE_DB_PASSWORD", &dbc.DbPass)
}

func loadEnvStr(key string, result *string) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	*result = s
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	n, err := strconv.Atoi(s)

	if err != nil {
		return
	}

	*result = uint(n) // will clamp the negative value
}
