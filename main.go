package main

import (
	"log"

	serverConstants "github.com/jettspanner123/AICodeWrapperBackend/constants"
	"github.com/joho/godotenv"

	gin "github.com/gin-gonic/gin"
	database "github.com/jettspanner123/AICodeWrapperBackend/db"
)

func main() {

	// MARK: Initializers
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error leading enviornment variables! 😭")
	}

	database.ConnectToDB()
	database.MigrateModels()

	// MARK: Application
	app := gin.Default()

	app.Run(serverConstants.GetPort())
}
