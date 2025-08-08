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
	seed.SeedLinear()
	seed.SeedLimit()	
	seed.SeedTurunan()
	seed.SeedIntegral()
}

func main() {
	router := gin.Default()

	Routes.Routes(router)

	router.Run(":" + os.Getenv("PORT"))
}
