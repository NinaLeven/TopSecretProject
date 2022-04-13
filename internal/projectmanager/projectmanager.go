package projectmanager

import "context"

type ProjectManagementService interface {
	CreateProject(ctx context.Context, db Storage, r ProjectCreateRequest) (*Project, error)
	UpdateProject(ctx context.Context, db Storage, r ProjectUpdateRequest) (*Project, error)
	GetProject(ctx context.Context, db Storage, r ProjectGetOptions) (*Project, error)
	ListProjects(ctx context.Context, db Storage, r ProjectListOptions) (*ProjectList, error)
}

type projectManagementService struct {
	employee EmployeeService
}

func NewProjectManagementService(employee EmployeeService) ProjectManagementService {
	return &projectManagementService{
		employee: employee,
	}
}
