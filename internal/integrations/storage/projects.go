package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"
	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

// internal storage model; in principal might be different to business-logic model; might be generated from sql
type project struct {
	UID            string    `sql:"uid"`
	Name           string    `sql:"name"`
	OwnerID        string    `sql:"owner_id"`
	State          string    `sql:"state"`
	ParticipantIDs []string  `sql:"participant_i_ds"`
	Progress       int32     `sql:"progress"`
	CreatedAt      time.Time `sql:"created_at"`
	UpdatedAt      time.Time `sql:"updated_at"`
}

const (
	projectsTable = "project"

	projectColumnUID            = "uid"
	projectColumnName           = "name"
	projectColumnOwnerID        = "owner_id"
	projectColumnState          = "state"
	projectColumnParticipantIDs = "participant_ids"
	projectColumnProgress       = "progress"
	projectColumnCreatedAt      = "created_at"
	projectColumnUpdatedAt      = "updated_at"
)

func (s *Repository) CreateProject(ctx context.Context, r projectmanager.StorageProjectCreateRequest) error {
	query, args, err := sq.Insert(projectsTable).
		Columns(
			projectColumnUID,
			projectColumnName,
			projectColumnOwnerID,
			projectColumnState,
			projectColumnParticipantIDs,
			projectColumnProgress,
			projectColumnCreatedAt,
			projectColumnUpdatedAt,
		).
		Values(
			r.UID,
			r.Name,
			r.OwnerID,
			r.State,
			pq.Array(r.ParticipantIDs),
			r.Progress,
			time.Now(),
			time.Now(),
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("unable to form query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("unable to insert project: %w", err)
	}

	return nil
}

func (s *Repository) UpdateProject(ctx context.Context, r projectmanager.StorageProjectUpdateRequest) error {
	queryBuilder := sq.Update(projectsTable).
		Where(sq.Eq{projectColumnUID: r.UID}).
		Set(projectColumnUpdatedAt, time.Now()).
		PlaceholderFormat(sq.Dollar)

	if r.Name != nil {
		queryBuilder = queryBuilder.Set(projectColumnName, *r.Name)
	}
	if r.State != nil {
		queryBuilder = queryBuilder.Set(projectColumnState, *r.State)
	}
	if r.ParticipantIDs != nil {
		queryBuilder = queryBuilder.Set(projectColumnParticipantIDs, pq.Array(*r.ParticipantIDs))
	}
	if r.Progress != nil {
		queryBuilder = queryBuilder.Set(projectColumnProgress, *r.Progress)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("unable to form query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("unable to update project: %w", err)
	}

	return nil
}

func (s *Repository) ListProjects(ctx context.Context, r projectmanager.ProjectListOptions) (*projectmanager.ProjectList, error) {
	queryBuilder := sq.Select(
		projectColumnUID,
		projectColumnName,
		projectColumnOwnerID,
		projectColumnState,
		projectColumnParticipantIDs,
		projectColumnProgress,
		projectColumnCreatedAt,
		projectColumnUpdatedAt,
	).
		From(projectsTable).
		PlaceholderFormat(sq.Dollar)

	if r.UIDs != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{projectColumnUID: *r.UIDs})
	}
	if r.Name != nil {
		queryBuilder = queryBuilder.Where(sq.ILike{projectColumnName: *r.Name})
	}
	if r.State != nil {
		queryBuilder = queryBuilder.Where(sq.ILike{projectColumnState: *r.State})
	}
	if r.Pagination != nil {
		queryBuilder = queryBuilder.Limit(uint64(r.Pagination.Limit)).
			Offset(uint64(r.Pagination.Offset))
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to form query: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("unable to select projects: %w", err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var res []project
	for rows.Next() {
		p := project{}
		arr := pq.StringArray{}
		err = rows.Scan(
			&p.UID,
			&p.Name,
			&p.OwnerID,
			&p.State,
			&arr,
			&p.Progress,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row project: %w", err)
		}
		p.ParticipantIDs = arr
		res = append(res, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("unable to scan projects: %w", err)
	}

	return &projectmanager.ProjectList{
		Projects: projectsToModel(res),
	}, nil
}

func projectsToModel(r []project) []projectmanager.Project {
	res := make([]projectmanager.Project, 0, len(r))
	for i := range r {
		res = append(res, *projectToModel(&r[i]))
	}
	return res
}

func projectToModel(r *project) *projectmanager.Project {
	return &projectmanager.Project{
		UID:            r.UID,
		Name:           r.Name,
		OwnerID:        r.OwnerID,
		State:          projectmanager.ProjectState(r.State),
		ParticipantIDs: r.ParticipantIDs,
		Progress:       r.Progress,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}
