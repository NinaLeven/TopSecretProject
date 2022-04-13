package projectmanager

import (
	"fmt"
	"time"
)

type ProjectState string

const (
	ProjectStatePlanned ProjectState = "planned"
	ProjectStateActive  ProjectState = "active"
	ProjectStateDone    ProjectState = "done"
	ProjectStateFailed  ProjectState = "failed"
)

func ProjectStateFromString(s string) (ProjectState, error) {
	switch s {
	case string(ProjectStatePlanned):
		return ProjectStatePlanned, nil
	case string(ProjectStateActive):
		return ProjectStateActive, nil
	case string(ProjectStateDone):
		return ProjectStateDone, nil
	case string(ProjectStateFailed):
		return ProjectStateFailed, nil
	default:
		return "", fmt.Errorf("invalid project state: %s", s)
	}
}

type ProjectCreateRequest struct {
	UID            string
	Name           string
	OwnerID        string
	State          *ProjectState
	Progress       *int32
	ParticipantIDs []string
}

type Project struct {
	UID            string
	Name           string
	OwnerID        string
	State          ProjectState
	ParticipantIDs []string
	Progress       int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ProjectUpdateRequest struct {
	UID            string
	Name           *string
	OwnerID        *string
	State          *ProjectState
	Progress       *int32
	ParticipantIDs *[]string
}

type ProjectGetOptions struct {
	UID string
}

type Pagination struct {
	Offset int
	Limit  int
}

type ProjectListOptions struct {
	UIDs       *[]string
	Name       *string
	State      *ProjectState
	Pagination *Pagination
}

type ProjectList struct {
	Projects []Project
	Hits     int
}
