package projectmanager

import "time"

type ProjectState string

const (
	ProjectStatePlanned ProjectState = "PLANNED"
	ProjectStateActive  ProjectState = "ACTIVE"
	ProjectStateDone    ProjectState = "DONE"
	ProjectStateFailed  ProjectState = "FAILED"
)

type ProjectCreateRequest struct {
	UID            string
	Name           string
	OwnerID        string
	State          *ProjectState
	Progress       *int
	ParticipantIDs []string
}

type Project struct {
	UID            string
	Name           string
	OwnerID        string
	State          ProjectState
	ParticipantIDs []string
	Progress       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ProjectUpdateRequest struct {
	UID            string
	Name           *string
	OwnerID        *string
	State          *ProjectState
	Progress       *int
	ParticipantIDs *[]string
}

type ProjectGetOptions struct {
	UID string
}

type Pagination struct {
	Offset int
	Length int
}

type PaginationWithHits struct {
	Pagination *Pagination
	Hits       int
}

type ProjectListOptions struct {
	UIDs       *[]string
	Name       *string
	Pagination *Pagination
}

type ProjectList struct {
	Projects   []Project
	Pagination PaginationWithHits
}
