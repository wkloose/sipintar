package main

import (
	"os"
	"github.com/wkloose/tempproject.git/initializers"
	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/Routes"
	"github.com/wkloose/tempproject.git/initializers/seed"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	seed.SeedTurunan()
	seed.SeedIntegral()
	seed.SeedLimit()
	seed.SeedLinear()
}

func main() {
	router := gin.Default()

	Routes.Routes(router)

	router.Run(":" + os.Getenv("PORT"))
}
