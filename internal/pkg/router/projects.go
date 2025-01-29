package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hse-revizor/projects-service/internal/pkg/router/dto"
	"github.com/hse-revizor/projects-service/internal/pkg/service/project"
)

// @Summary Create project
// @Description In success case returns created project model.
// @Tags Project
// @Accept json
// @Param data body dto.CreateProjectDto true "Project input"
// @Success 200 "" ""
// @Router /projects [post]
func (h *Handler) CreateProject(c *gin.Context) {
	var req dto.ProjectDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := h.service.CreateProject(c, &project.CreateProject{
		Name:          req.Name,
		RepositoryURL: req.RepositoryURL,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, project)
}

// @Summary Get project by id
// @Description In success case returns project model with provided id
// @Tags Project
// @Param id path string true "Project id input"
// @Success 200 "" ""
// @Router /project/{id} [get]
func (h *Handler) GetProject(c *gin.Context) {
	id := c.Param("id")
	projectUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	project, err := h.service.GetProjectById(c, projectUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	res := dto.GetProjectDto{
		Name:          project.Name,
		LastCheckDate: "",
		RepositoryURL: project.Sources[0],
		FileName:      "",
	}
	responseOK(c, project)
}

// @Summary Delete project by id
// @Description In success case delete project model with provided id
// @Tags Project
// @Param id path string true "Project id input"
// @Success 200 "" ""
// @Router /project/{id} [delete]
func (h *Handler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	projectUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	project, err := h.service.DeleteProject(c, projectUUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	responseOK(c, project)
}
