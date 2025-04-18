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

	r.POST("/login", handlers.LoginHandler)

	//listeners
	r.POST("/listeners/create", handlers.CreateListener)
	r.PUT("/listeners/:id", handlers.UpdateListenerData)
	r.DELETE("/listeners/:id", handlers.DeleteListener)
	r.POST("/listeners/", handlers.ReadListener)

	//AboutListener
	r.GET("/listeners/:id/about", handlers.AboutListener)

	//Enrollments
	r.GET("/listeners/:id/about/enrollments", handlers.GetProgramInfo)
	r.GET("/listeners/:id/about/enrollments/program", handlers.SelectProgramEducation)

	//
	r.GET("/levelEducation", handlers.GetAllLevelsEducation)

	r.Run("localhost:8080")

}
