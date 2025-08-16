package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/jandrey/run-project-go/src/controller/user"
)

func InitRoutes(r *gin.RouterGroup) {
	user := r.Group("/users") // base RESTful

	user.GET("/:id", controller.FindUserById)
	user.GET("/email/:email", controller.FindUserByEmail)
	user.POST("/", controller.CreateUser)
	user.PUT("/:id", controller.UpdateUser)
	user.DELETE("/:id", controller.DeleteUser)
}
