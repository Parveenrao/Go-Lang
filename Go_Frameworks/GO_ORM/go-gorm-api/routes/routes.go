package routes

import (
	"github.com/gin-gonic/gin"
	"go-gorm-api/controllers"
)

func SetupRoute(r *gin.Engine) {
	r.POST("/users" , controllers.CreatUser)
	r.GET("/users" , controllers.GetUsers)
}