package handlers

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
		auth.POST("/logout", h.logout)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getList)
			lists.POST("/", h.createList)
			lists.GET("/:list_id", h.getListDetail)
			lists.PUT("/:list_id", h.updateList)
			lists.DELETE("/:list_id", h.deleteList)
			tasks := lists.Group("/tasks")
			{
				tasks.GET("/", h.getTasks)
				tasks.POST("/", h.createTask)
				tasks.GET("/:task_id", h.getTaskDetail)
				tasks.PUT("/:task_id", h.updateTask)
				tasks.DELETE("/:task_id", h.deleteTask)
			}
		}
	}
	return router
}
