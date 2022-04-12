package api

import (
	"net/http"

	swagger "github.com/NinaLeven/TopSecretProject/api/generated"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type ProjectListResponse struct {
	Projects []swagger.ProjectResponse `json:"projects"`
	Hits     int                       `json:"hits"`
}

func (p *projectManagerServer) listProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// may not use transactions for simple selects
	res, err := p.projectManagementService.ListProjects(ctx,
		p.storageFactory.NewRepository(),
		projectmanager.ProjectListOptions{
			Name:       nil,
			Pagination: nil,
		},
	)
	if err != nil {
		p.log.WithError(err).Error("unable to list project")
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, ProjectListResponse{
		Projects: modelToProjectListResponse(res.Projects),
		Hits:     res.Hits,
	})
}

func modelToProjectListResponse(r []projectmanager.Project) []swagger.ProjectResponse {
	res := make([]swagger.ProjectResponse, 0, len(r))
	for i := range r {
		res = append(res, *modelToProjectResponse(&r[i]))
	}
	return res
}
