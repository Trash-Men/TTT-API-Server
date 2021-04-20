package main // import "github.com/Trash-Men/api-server"

import (
	"github.com/Trash-Men/api-server/api/routes"
	"github.com/Trash-Men/api-server/configs"
	"github.com/Trash-Men/api-server/database"
	"github.com/Trash-Men/api-server/repositories"
)

func main() {
	database.PostgreSQLConnection()

	defer database.DBClient.Close()

	repositories.SetDBClient(database.DBClient)

	router := routes.SetRouting()

	if configs.GetEnvironments().MODE == "development" {
		router.Static("/upload", "./upload")
	}

	router.Logger.Fatal(router.Start(":" + "5000"))
}
