package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"main/server/handlers/educationHandlers"
	"main/server/handlers/enrollmentHandlers"
	"main/server/handlers/listenerHandlers"
	"main/server/handlers/loginHandler"
	"main/server/handlers/programHandlers"
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

	r.POST("/login", loginHandler.LoginHandler)

	//listeners
	r.POST("/listeners/create", listenerHandlers.CreateListener)
	r.PUT("/listeners/update/:id", listenerHandlers.UpdateListener)
	r.DELETE("/listeners/:id", listenerHandlers.DeleteListener)
	r.POST("/listeners/", listenerHandlers.GetListeners)
	r.GET("/listeners/:id/about", listenerHandlers.GetListenerByID)

	//Enrollments
	r.GET("/listeners/:id/about/enrollments/program", enrollmentHandlers.SelectProgramEducation)
	r.POST("listeners/:id/about/enrollments/program/record", enrollmentHandlers.RecordListenerOnProgram)

	//LevelEducation
	r.GET("/levelEducation", educationHandlers.GetAllLevelsEducation)

	//DivisionsEducation
	r.GET("/divisionsEducation", educationHandlers.GetAllDivisions)

	//Educationtype
	r.GET("/educationtypes", educationHandlers.GetAllEducationTypes)

	//ProgramEducation
	r.POST("/programeducation/create", programHandlers.CreateProgram)
	r.DELETE("/programeducation/:id/delete", programHandlers.DeleteProgram)
	r.PUT("/programeducation/:id/update", programHandlers.UpdateProgram)
	r.POST("/programeducation", programHandlers.GetPrograms)

	//documents

	r.Run("localhost:8080")

}
