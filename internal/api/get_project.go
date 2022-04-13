package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	swagger "github.com/NinaLeven/TopSecretProject/api/generated"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

func modelToProjectResponse(r *projectmanager.Project) *swagger.ProjectResponse {
	return &swagger.ProjectResponse{
		Uid:            r.UID,
		Name:           r.Name,
		OwnerId:        r.OwnerID,
		State:          string(r.State),
		Progress:       r.Progress,
		ParticipantIds: r.ParticipantIDs,
		CreatedAt:      r.CreatedAt.String(),
		UpdatedAt:      r.UpdatedAt.String(),
	}
}

func (p *projectManagerServer) getProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uid := chi.URLParam(r, "uid")

	var res *projectmanager.Project

	// may not use transactions for simple selects
	res, err := p.projectManagementService.GetProject(ctx,
		p.storageFactory.NewRepository(),
		projectmanager.ProjectGetOptions{UID: uid},
	)
	if err != nil {
		switch err := err.(type) {
		case *projectmanager.ProjectNotFound:
			respondWithJSON(w, http.StatusNotFound, err)
			return
		default:
			p.log.WithError(err).Error("unable to get project")
			respondWithJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, modelToProjectResponse(res))
}
