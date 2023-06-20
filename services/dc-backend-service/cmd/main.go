package main

import (
	"context"
	"dc-backend/internal/api"
	"dc-backend/internal/storage"
	"dc-backend/internal/token"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	PGUSER           = "PGUSER"
	PGPASSWORD       = "PGPASSWORD"
	PGHOST           = "PGHOST"
	PGPORT           = "PGPORT"
	PGDATABASE       = "PGDATABASE"
	JWTSECRET        = "JWTSECRET"
	DCBACKENDAPIPORT = "DCBACKENDAPIPORT"
)

type Config struct {
	dbUser       string
	dbPassword   string
	dbHost       string
	dbPort       string
	dbDatabase   string
	jwtSecretKey string
	apiPort      string
}

func main() {
	config := readEnvVars()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.dbUser,
		config.dbPassword,
		config.dbHost,
		config.dbPort,
		config.dbDatabase)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	jwtValidator := token.NewJWTValidator(config.jwtSecretKey)
	itemStorage := storage.NewItemStorage(conn)
	server := api.New(fmt.Sprintf(":%s", config.apiPort), jwtValidator, itemStorage)
	err = server.Run()
	if err != nil {
		log.Fatalf("cannot start server: %v\n", err)
	}
}

func readEnvVars() Config {
	dbUser := readEnvVar(PGUSER)
	dbPassword := readEnvVar(PGPASSWORD)
	dbHost := readEnvVar(PGHOST)
	dbPort := readEnvVar(PGPORT)
	dbDatabase := readEnvVar(PGDATABASE)
	jwtSecretKey := readEnvVar(JWTSECRET)
	apiPort := readEnvVar(DCBACKENDAPIPORT)

	return Config{
		dbUser:       dbUser,
		dbPassword:   dbPassword,
		dbHost:       dbHost,
		dbPort:       dbPort,
		dbDatabase:   dbDatabase,
		jwtSecretKey: jwtSecretKey,
		apiPort:      apiPort,
	}
}

func readEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("%s env var is not set", name)
	}

	return value
}
