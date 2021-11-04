package main

import (
	"log"
	"os"
	"schapi/database"
	"schapi/endpoint/authapi"
	"schapi/endpoint/scheduleapi"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if gin.Mode() == gin.DebugMode {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	if gin.Mode() == gin.DebugMode {
		config.AllowOrigins = []string{"http://localhost:8080"}
	} else {
		config.AllowOrigins = []string{"https://vue-sch.web.app/"}
	}

	database.Initialize()

	router.Use(cors.New(config))

	// Api
	api := router.Group("/api")
	{
		authapi.SetupRoute(api)

		// 認証必要なエンドポイント
		api.Use(authapi.MiddlewareFunc())
		{
			authapi.SetupAuthenticatedRoute(api)
			scheduleapi.SetupRoute(api)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}
