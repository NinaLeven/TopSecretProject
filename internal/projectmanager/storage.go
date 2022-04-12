package projectmanager

import (
	"context"
)

type StorageProjectCreateRequest struct {
	UID            string
	Name           string
	OwnerID        string
	State          ProjectState
	Progress       int
	ParticipantIDs []string
}

type StorageProjectUpdateRequest struct {
	UID            string
	Name           *string
	OwnerID        *string
	State          *ProjectState
	Progress       *int
	ParticipantIDs *[]string
}

type Storage interface {
	CreateProject(ctx context.Context, r StorageProjectCreateRequest) error
	UpdateProject(ctx context.Context, r StorageProjectUpdateRequest) error
	ListProjects(ctx context.Context, r ProjectListOptions) (*ProjectList, error)
}
