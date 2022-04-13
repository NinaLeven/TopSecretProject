package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

// UpdateProjectRequest not using generated one since those lack nullability
type UpdateProjectRequest struct {
	// Project name
	Name *string `json:"name,omitempty"`
	// Project owner id
	OwnerId *string `json:"owner_id,omitempty"`
	// Project state; Might be created non-delault for creating prioject post-factum
	State *string `json:"state,omitempty"`
	// Project progress in %
	Progress *int32 `json:"progress,omitempty"`
	// Ids of the participants
	ParticipantIds *[]string `json:"participant_ids,omitempty"`
}

func validateUpdateProject(project *UpdateProjectRequest) error {
	if project.Name != nil && *project.Name == "" {
		return fmt.Errorf("name must not be empty")
	}
	if project.OwnerId != nil && *project.OwnerId == "" {
		return fmt.Errorf("owner_id must not be empty")
	}
	if project.Progress != nil && (*project.Progress < 0 || *project.Progress > 100) {
		return fmt.Errorf("progress is out of bounds")
	}
	if project.State != nil {
		_, err := projectmanager.ProjectStateFromString(*project.State)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *projectManagerServer) updateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uid := chi.URLParam(r, "uid")

	// not using generated one since those lack nullability
	project := UpdateProjectRequest{}

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		p.log.WithError(err).Error("unable to decode request")
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("unable to decode request: %s", err))
		return
	}

	err = validateUpdateProject(&project)
	if err != nil {
		p.log.WithError(err).Error("unable to decode request")
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}

	var res *projectmanager.Project

	err = p.storageFactory.RunInTransaction(ctx, func(ctx context.Context, storage projectmanager.Storage) error {
		res, err = p.projectManagementService.UpdateProject(ctx, storage,
			projectmanager.ProjectUpdateRequest{
				UID:            uid,
				Name:           project.Name,
				OwnerID:        project.OwnerId,
				State:          (*projectmanager.ProjectState)(project.State),
				Progress:       project.Progress,
				ParticipantIDs: project.ParticipantIds,
			},
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		switch err := err.(type) {
		case *projectmanager.OwnerRoleError, *projectmanager.ParticipantDepartmentError:
			respondWithJSON(w, http.StatusBadRequest, err)
			return
		case *projectmanager.ProjectNotFound:
			respondWithJSON(w, http.StatusNotFound, err)
			return
		default:
			p.log.WithError(err).Error("unable to update project")
			respondWithJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, modelToProjectResponse(res))
}
