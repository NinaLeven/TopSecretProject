package projectmanager

import (
	"context"
	"fmt"
)

type ParticipantDepartmentError struct {
	ParticipantID         string
	ParticipantDepartment string
	OwnerDepartment       string
}

func (p *ParticipantDepartmentError) Error() string {
	return fmt.Sprintf("participant %s is from the different department from owner: expected: %s, got: %s",
		p.ParticipantID,
		p.OwnerDepartment,
		p.ParticipantDepartment,
	)
}

func (p *projectManagementService) checkParticipantsDepartments(ctx context.Context, participantIDs []string, department string) error {
	for i := range participantIDs {
		employee, err := p.employee.GetEmployee(ctx, participantIDs[i])
		if err != nil {
			return fmt.Errorf("unable to get participant: %w", err)
		}

		if employee.Department != department {
			return &ParticipantDepartmentError{
				ParticipantID:         employee.Id,
				ParticipantDepartment: employee.Department,
				OwnerDepartment:       department,
			}
		}
	}
	return nil
}

type OwnerRoleError struct {
	OwnerID      string
	OwnerRole    EmployeeRole
	ExpectedRole EmployeeRole
}

func (o *OwnerRoleError) Error() string {
	return fmt.Sprintf("incorrect owner(%s) role: expected: %s, got: %s", o.OwnerID, o.ExpectedRole, o.OwnerRole)
}

func (p *projectManagementService) checkOwner(owner *Employee) error {
	if owner.Role != EmployeeRoleManager {
		return &OwnerRoleError{
			OwnerID:      owner.Id,
			OwnerRole:    owner.Role,
			ExpectedRole: EmployeeRoleManager,
		}
	}

	return nil
}

func (p *projectManagementService) CreateProject(ctx context.Context, db Storage, r ProjectCreateRequest) (
	*Project,
	error,
) {
	owner, err := p.employee.GetEmployee(ctx, r.OwnerID)
	if err != nil {
		return nil, fmt.Errorf("unable to get owner %s: %w", r.OwnerID, err)
	}

	err = p.checkOwner(owner)
	if err != nil {
		return nil, err
	}

	err = p.checkParticipantsDepartments(ctx, r.ParticipantIDs, owner.Department)
	if err != nil {
		return nil, err
	}

	req := StorageProjectCreateRequest{
		UID:            r.UID,
		Name:           r.Name,
		OwnerID:        r.OwnerID,
		ParticipantIDs: r.ParticipantIDs,
		State:          ProjectStatePlanned,
		Progress:       0,
	}
	if r.State != nil {
		req.State = *r.State
	}
	if r.Progress != nil {
		req.State = *r.State
	}

	err = db.CreateProject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to create project: %w", err)
	}

	return p.GetProject(ctx, db, ProjectGetOptions{
		UID: r.UID,
	})
}

func (p *projectManagementService) UpdateProject(ctx context.Context, db Storage, r ProjectUpdateRequest) (
	*Project,
	error,
) {
	project, err := p.GetProject(ctx, db, ProjectGetOptions{
		UID: r.UID,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get current project: %w", err)
	}

	var owner *Employee

	if r.OwnerID != nil {
		owner, err = p.employee.GetEmployee(ctx, *r.OwnerID)
		if err != nil {
			return nil, fmt.Errorf("unable to get owner %s: %w", r.OwnerID, err)
		}

		err = p.checkOwner(owner)
		if err != nil {
			return nil, err
		}
	}

	if r.ParticipantIDs != nil {
		if owner != nil {
			owner, err = p.employee.GetEmployee(ctx, project.OwnerID)
			if err != nil {
				return nil, fmt.Errorf("unable to get owner %s: %w", r.OwnerID, err)
			}
		}

		err = p.checkParticipantsDepartments(ctx, *r.ParticipantIDs, owner.Department)
		if err != nil {
			return nil, err
		}
	}

	req := StorageProjectUpdateRequest{
		UID:            r.UID,
		Name:           r.Name,
		OwnerID:        r.OwnerID,
		State:          r.State,
		Progress:       r.Progress,
		ParticipantIDs: r.ParticipantIDs,
	}

	err = db.UpdateProject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to update project: %w", err)
	}

	return p.GetProject(ctx, db, ProjectGetOptions{
		UID: r.UID,
	})
}

type ProjectNotFound struct {
	UID string
}

func (p *ProjectNotFound) Error() string {
	return fmt.Sprintf("project(%s) not found", p.UID)
}

func (p *projectManagementService) GetProject(ctx context.Context, db Storage, r ProjectGetOptions) (*Project, error) {
	projects, err := p.ListProjects(ctx, db, ProjectListOptions{
		UIDs: &[]string{r.UID},
	})
	if err != nil {
		return nil, fmt.Errorf("umnable to get project: %w", err)
	}
	if len(projects.Projects) < 1 {
		return nil, &ProjectNotFound{UID: r.UID}
	}

	return &projects.Projects[0], nil
}

func (p *projectManagementService) ListProjects(ctx context.Context, db Storage, r ProjectListOptions) (*ProjectList, error) {
	return db.ListProjects(ctx, r)
}
