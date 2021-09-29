package main

import (
	"log"

	"pelatihan-be/database"

	util "pelatihan-be/helpers/utils"
	"pelatihan-be/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))

}

func SetupRouter() *gin.Engine {

	db := database.Connection()

	r := gin.Default()

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	routes.InitAuthRoutes(db, r)
	routes.InitTokenRoutes(db, r)

	return r

}
