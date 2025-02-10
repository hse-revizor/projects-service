package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hse-revizor/projects-service/internal/pkg/router/dto"
	"github.com/hse-revizor/projects-service/internal/pkg/service/project"
)

// @Summary Create project
// @Description Create a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param data body dto.ProjectDto true "Project input"
// @Success 200 {object} dto.GetProjectDto
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /project [post]
func (h *Handler) CreateProject(c *gin.Context) {
	var req dto.ProjectDto
	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
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
	res := dto.GetProjectDto{
		Id:            project.Id.String(),
		Name:          project.Name,
		LastCheckDate: "",
		RepositoryURL: project.Sources[0],
	}
	responseOK(c, res)
}

// @Summary Get project by id
// @Description Get project details by its ID
// @Tags Project
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200 {object} dto.GetProjectDto
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /project/{projectId} [get]
func (h *Handler) GetProject(c *gin.Context) {
	id := c.Param("projectId")
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
	lastCheckDate := ""
	if project.LastCheckDate != nil {
		lastCheckDate = project.LastCheckDate.Format(time.RFC3339)
	}
	res := dto.GetProjectDto{
		Id:            project.Id.String(),
		Name:          project.Name,
		LastCheckDate: lastCheckDate,
		RepositoryURL: project.Sources[0],
	}
	responseOK(c, res)
}

// @Summary Delete project by id
// @Description Delete a project by its ID
// @Tags Project
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200 {object} dto.GetProjectDto
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /project/{projectId} [delete]
func (h *Handler) DeleteProject(c *gin.Context) {
	id := c.Param("projectId")
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
	res := dto.GetProjectDto{
		Id:            project.Id.String(),
		Name:          project.Name,
		LastCheckDate: "",
		RepositoryURL: project.Sources[0],
	}
	responseOK(c, res)
}
