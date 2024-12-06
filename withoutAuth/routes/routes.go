package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"withoutAuth/controller"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/students", controller.GetAllStudent)
	r.POST("/students", controller.CreateStudent)
	r.GET("/students/id/:id", controller.GetStudentById)
	r.GET("/students/name/:name", controller.GetIdByStudent)
	r.PATCH("/students/:id", controller.UpdateStudent)
	r.DELETE("students/:id", controller.DeleteStudent)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
