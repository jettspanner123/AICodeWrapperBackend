package main

import (
	"log"

	serverConstants "github.com/jettspanner123/AICodeWrapperBackend/constants"
	"github.com/joho/godotenv"

	gin "github.com/gin-gonic/gin"
	database "github.com/jettspanner123/AICodeWrapperBackend/db"

	routes "github.com/jettspanner123/AICodeWrapperBackend/routes"
)

func main() {

	// MARK: Initializers
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error leading enviornment variables! 😭")
	}

	// MARK: Database Init & Migration
	database.ConnectToDB()
	database.MigrateModels()

	// MARK: Application
	app := gin.Default()

	// MARK: Clustering Routes
	routes.RegisterAuthRoutes(app)

	app.Run(serverConstants.GetPort())
}
