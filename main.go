package main

import (
	"log"
	"os"
	"schapi/database"
	"schapi/endpoint/authapi"
	"schapi/endpoint/scheduleapi"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	config.AllowOrigins = []string{
		"http://localhost:8080",
		"https://vue-sch.web.app",
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
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}
