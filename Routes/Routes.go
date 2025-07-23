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
	}

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth)
	{

	}
}

