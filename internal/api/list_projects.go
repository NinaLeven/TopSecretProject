package api

import (
	"net/http"
	"net/url"
	"strconv"

	swagger "github.com/NinaLeven/TopSecretProject/api/generated"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type ProjectListResponse struct {
	Projects []swagger.ProjectResponse `json:"projects"`
	Hits     int                       `json:"hits"`
}

func paginationFromQuery(query url.Values) *projectmanager.Pagination {
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")

	if limitStr == "" || offsetStr == "" {
		return nil
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return nil
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return nil
	}

	return &projectmanager.Pagination{
		Offset: offset,
		Limit:  limit,
	}
}

func (p *projectManagerServer) listProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := r.URL.Query()

	opts := projectmanager.ProjectListOptions{
		Pagination: paginationFromQuery(query),
	}

	name := query.Get("name")
	if name != "" {
		opts.Name = &name
	}

	state := query.Get("state")
	if state != "" {
		s, err := projectmanager.ProjectStateFromString(state)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, err)
			return
		}
		opts.State = &s
	}

	// may not use transactions for simple selects
	res, err := p.projectManagementService.ListProjects(ctx, p.storageFactory.NewRepository(), opts)
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
