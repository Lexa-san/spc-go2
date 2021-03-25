package routers

import (
	"github.com/Lexa-san/spc-go2/12.GinGormAPI/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // mux.NewRouter() analog

	//	Gin-way is to group resources
	apiV1Group := router.Group("/api/v1") //prefix
	//imitate function execution
	{
		apiV1Group.GET("article", handlers.GetAllArticles)
		apiV1Group.POST("article", handlers.PostNewArticle)
		apiV1Group.GET("article/:id", handlers.GetArticleById)
		apiV1Group.PUT("article/:id", handlers.UpdateArticleById)
		apiV1Group.DELETE("article/:id", handlers.DeleteArticleById)
	}

	return router
}
