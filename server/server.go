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
	r.PUT("/listeners/:id", handlers.UpdateListener)
	r.DELETE("/listeners/:id", handlers.DeleteListener)
	r.POST("/listeners/", handlers.ReadListener)

	// //LevelEducation
	// r.PUT("/levelEducation/:id", handlers.LevelEducationUpdate)
	// r.POST("/levelEducation", handlers.LevelEducationCreate)

	//AboutListener
	r.GET("/listeners/:id/about", handlers.AboutListener)
	r.PUT("/listeners/:id/about/update/passport", handlers.UpdateListenersPassport)
	r.PUT("/listeners/:id/about/update/education", handlers.UpdateListenersPassport)
	r.PUT("/listeners/:id/about/update/placework", handlers.UpdateListenersPlaceWork)
	r.PUT("/listeners/:id/about/update/programeducation", handlers.UpdateListenersPassport)
	r.PUT("/listeners/:id/about/update/regaddress", handlers.UpdateListenersRegAddress)

	r.Run("localhost:8080")

}
