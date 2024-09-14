package router

import (
	"code/app/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		//user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		auth := api.Group("/auth")
		auth.POST("/login", init.AuthCtrl.Login)
		auth.POST("/forget", init.AuthCtrl.Forget)

		//auth.GET("/:signup", init.A.GetUserById)
		//auth.PUT("/:userID", init.UserCtrl.UpdateUserData)
	}

	return router
}
