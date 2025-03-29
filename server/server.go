package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"main/server/handlers"
)

func Start() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/login", handlers.PostUserHandler)

	//listeners
	r.POST("/listeners", handlers.CreateListener)
	r.PUT("/listeners/:id", handlers.UpdateListener)
	r.DELETE("/listeners/:id", handlers.DeleteListener)
	r.GET("/listeners", handlers.ReadListener)

	// //LevelEducation
	// r.PUT("/levelEducation/:id", handlers.LevelEducationUpdate)
	// r.POST("/levelEducation", handlers.LevelEducationCreate)
	r.Run("localhost:8080")

}
