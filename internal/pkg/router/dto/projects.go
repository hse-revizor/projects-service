package dto

type ProjectDto struct {
	Name          string `json:"name" binding:"required"`
	RepositoryURL string `json:"repositoryURL" binding:"required"`
} // @name ProjectDto

type GetProjectDto struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LastCheckDate string `json:"lastCheckDate"`
	RepositoryURL string `json:"repositoryURL"`
} // @name GetProjectDto

type ProjectGroupDto struct {
	Name  string   `json:"name" binding:"required"`
	Items []string `json:"items"`
} // @name ProjectGroupDto
