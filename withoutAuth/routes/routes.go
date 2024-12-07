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
	r.GET("/tugaspendahuluans", controller.GetAllTugasPendahuluan)
	r.POST("/tugaspendahuluans", controller.CreateTugasPendahuluan)
	r.GET("/tugaspendahuluans/id/:id", controller.GetTugasPendahuluanById)
	r.PATCH("/tugaspendahuluans/:id", controller.UpdateTugasPendahuluan)
	r.DELETE("/tugaspendahuluans/:id", controller.DeleteTugasPendahuluan)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
