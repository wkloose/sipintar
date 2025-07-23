package main

import (
	"os"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/Routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	Routes.Routes(router)

	router.Run(":" + os.Getenv("PORT"))
}
