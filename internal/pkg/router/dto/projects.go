package dto

type ProjectDto struct {
    Name           string `json:"name" binding:"required"`
    RepositoryURL  string `json:"repositoryURL" binding:"required"`
}

type GetProjectDto struct {
    Name           string `json:"name"`
    LastCheckDate  string `json:"lastCheckDate"`
    RepositoryURL  string `json:"repositoryURL"`
    FileName       string `json:"fileName"`
}

type ProjectGroupDto struct {
    Name          string   `json:"name" binding:"required"`
    Items         []string `json:"items"`
}