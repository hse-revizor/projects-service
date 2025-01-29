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
	docs.SwaggerInfo.BasePath = "/api"
	router := api.Group("/api")
	{
		projects := router.Group("/project")
		{
			projects.POST("", h.CreateProject)
			projects.GET("/:projectId", h.GetProject)
			projects.DELETE("/:projectId", h.DeleteProject)
			group := projects.Group("/group")
			{
				group.POST("", h.CreateGroup)
				group.GET("/:groupId", h.GetGroup)
				group.DELETE("/:groupId", h.DeleteGroup)
			}
		}

	}
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.Use(cors.Default())

	return api
}
