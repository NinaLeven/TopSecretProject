package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

// CrateProjectRequest not using generated one since those lack nullability
type CrateProjectRequest struct {
	// Unique project identifier, might be used for idempotency
	Uid string `json:"uid"`
	// Project name
	Name string `json:"name"`
	// Project owner id
	OwnerId string `json:"owner_id"`
	// Project state; Might be created non-delault for creating prioject post-factum
	State *string `json:"state,omitempty"`
	// Project progress in %
	Progress *int32 `json:"progress,omitempty"`
	// Ids of the participants
	ParticipantIds []string `json:"participant_ids,omitempty"`
}

func validateCreateProject(project *CrateProjectRequest) error {
	if project.Uid == "" {
		return fmt.Errorf("project uid must be set")
	}
	if project.Name == "" {
		return fmt.Errorf("project name must be set")
	}
	if project.OwnerId == "" {
		return fmt.Errorf("project owner_id must be set")
	}
	if project.State != nil {
		_, err := projectmanager.ProjectStateFromString(*project.State)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *projectManagerServer) createProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// not using generated one since those lack nullability
	project := CrateProjectRequest{}

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		p.log.WithError(err).Error("unable to decode request")
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("unable to decode request: %s", err))
		return
	}

	err = validateCreateProject(&project)
	if err != nil {
		p.log.WithError(err).Error("unable to decode request")
		respondWithJSON(w, http.StatusBadRequest, err)
		return
	}

	var res *projectmanager.Project

	err = p.storageFactory.RunInTransaction(ctx, func(ctx context.Context, storage projectmanager.Storage) error {
		res, err = p.projectManagementService.CreateProject(ctx, storage,
			projectmanager.ProjectCreateRequest{
				UID:            project.Uid,
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
		default:
			p.log.WithError(err).Error("unable to create project")
			respondWithJSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, modelToProjectResponse(res))
}
