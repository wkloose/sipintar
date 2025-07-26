package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/controllers"
	"github.com/wkloose/tempproject.git/middleware"
)

func Routes(r *gin.Engine) {
	public := r.Group("/")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
		public.POST("/password/forgot", controllers.ForgotPassword)
		public.POST("/password/reset", controllers.ResetPassword)
	}

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth)
	{
	heart := protected.Group("/hearts")
	{
		heart.GET("/", controllers.GetHeartCount)
		heart.GET("/status", controllers.GetHeartStatus)
		heart.POST("/claim", controllers.ClaimHeart)
	}
	}

}

