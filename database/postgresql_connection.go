package database

import (
	"context"
	"fmt"

	"github.com/Trash-Men/api-server/configs"
	"github.com/go-pg/pg/v10"
)

var DBClient *pg.DB

func PostgreSQLConnection() {
	environments := configs.GetEnvironments()

	db := pg.Connect(&pg.Options{
		Addr:     environments.DB_HOST + ":" + environments.DB_PORT,
		User:     environments.DB_USER,
		Password: environments.DB_PASSWORD,
		Database: environments.DB_NAME,
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("connection...")

	DBClient = db

	fmt.Println("Successfully connected!")
}
