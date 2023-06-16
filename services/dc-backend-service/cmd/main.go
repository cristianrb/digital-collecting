package main

import (
	"context"
	"dc-backend/internal/api"
	"dc-backend/internal/storage"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	dbUrl := "postgresql://dc_admin:dc_password@localhost:5432/digital_collecting?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	itemStorage := storage.NewItemStorage(conn)
	
	server := api.New(":8080", itemStorage)
	err = server.Run()
	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
