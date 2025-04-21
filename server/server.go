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
	r.POST("listeners/:id/about/enrollments/program/record", handlers.RecordListenerOnProgram)

	//LevelEducation
	r.GET("/levelEducation", handlers.GetAllLevelsEducation)

	//DivisionsEducation
	r.GET("/divisionsEducation", handlers.GetAllDivisions)

	//Educationtype
	r.GET("/educationtypes", handlers.GetAllEducationTypes)

	//ProgramEducation
	r.POST("/programeducation/create", handlers.CreateProgram)
	r.DELETE("/programeducation/:id/delete", handlers.DeleteProgram)
	r.PUT("/programeducation/:id/update", handlers.UpdateProgram)
	r.POST("/programeducation", handlers.ReadProgram)

	r.Run("localhost:8080")

}
