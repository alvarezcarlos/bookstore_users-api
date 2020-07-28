package app

import (
	"github.com/alvarezcarlos/bookstore_users-api/controllers/ping"
	"github.com/alvarezcarlos/bookstore_users-api/controllers/users"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users", users.FindByStatus)
	router.POST("/users", users.CreateUser)
}