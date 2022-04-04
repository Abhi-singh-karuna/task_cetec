package main

import (
	"log"

	"github.com/Abhi-singh-karuna/config"
	"github.com/Abhi-singh-karuna/router"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// initilize the gin
	app := gin.Default()

	// stablish connection with database
	config.ConnectDB()

	// load the environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Routes(app)

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	log.Fatal(app.Run(":8000"))
}
