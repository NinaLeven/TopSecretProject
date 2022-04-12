package projectmanager

import "context"

type EmployeeRole string

const (
	EmployeeRoleManager  EmployeeRole = "manager"
	EmployeeRoleEmployee EmployeeRole = "employee"
)

type Employee struct {
	Id         string
	FirstName  string
	LastName   string
	Email      string
	Department string
	Role       EmployeeRole
}

type EmployeeService interface {
	GetEmployee(ctx context.Context, id string) (*Employee, error)
	ListEmployees(ctx context.Context) ([]Employee, error)
}
