package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/hse-revizor/projects-service/docs"
	"github.com/hse-revizor/projects-service/internal/pkg/service/project"
	"github.com/hse-revizor/projects-service/internal/utils/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	cfg     *config.Config
	service *project.Service
}

func NewRouter(cfg *config.Config, service *project.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	api := gin.New()

	api.Use(gin.Recovery())
	api.Use(gin.Logger())
	api.Use(cors.Default())

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.Title = "Projects Service API"
	docs.SwaggerInfo.Description = "API Documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8787"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router := api.Group("/api")
	{
		router.POST("/project", h.CreateProject)
		router.GET("/project/:projectId", h.GetProject)
		router.DELETE("/project/:projectId", h.DeleteProject)
		// group := router.Group("/group")
		// {
		// 	group.POST("", h.CreateGroup)
		// 	group.GET("/:groupId", h.GetGroup)
		// 	group.DELETE("/:groupId", h.DeleteGroup)
		// }
	}

	return api
}
