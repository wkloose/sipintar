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
	materials := protected.Group("/materials")
	{
		materials.GET("/", controllers.GetAllMaterials)
		materials.GET("/:id", controllers.GetMaterialDetail)
	}
	quizzes := protected.Group("/quiz")
	{
		quizzes.GET("/:material_id", controllers.GetQuizHandler)
		quizzes.POST("/:material_id/submit", controllers.SubmitQuizHandler)
	}
	roadmap := protected.Group("/roadmap")
	{
		roadmap.GET("/", controllers.GetRoadmap)
		roadmap.GET("/progress", controllers.GetTotalProgressHandler)
	}
	streak := protected.Group("/streak")
	{
		streak.GET("/", controllers.GetStreakHandler)
	}
	profile := protected.Group("/profile")
	{
		profile.GET("/", controllers.GetProfile)
		profile.PUT("/", controllers.UpdateProfile)
	}
	protected := r.Group("/progress")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("/statistik", controllers.GetProgressStatHandler)
	}
}
}