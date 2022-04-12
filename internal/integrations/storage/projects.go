package storage

import (
	"context"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

func (s *Repository) CreateProject(ctx context.Context, r projectmanager.StorageProjectCreateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *Repository) UpdateProject(ctx context.Context, r projectmanager.StorageProjectUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *Repository) ListProjects(ctx context.Context, r projectmanager.ProjectListOptions) (*projectmanager.ProjectList, error) {
	//TODO implement me
	panic("implement me")
}

