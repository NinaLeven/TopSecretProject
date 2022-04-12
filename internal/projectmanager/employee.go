package projectmanager

import (
	"context"
	"fmt"
)

type EmployeeRole string

const (
	EmployeeRoleManager  EmployeeRole = "manager"
	EmployeeRoleEmployee EmployeeRole = "employee"
)

func EmployeeRoleFromString(s string) (EmployeeRole, error) {
	switch s {
	case string(EmployeeRoleManager):
		return EmployeeRoleManager, nil
	case string(EmployeeRoleEmployee):
		return EmployeeRoleEmployee, nil
	default:
		return "", fmt.Errorf("unknown employee role: %s", s)
	}
}

type Employee struct {
	Id         string
	FirstName  string
	LastName   stringar
	Email      string
	Department string
	Role       EmployeeRole
}

type EmployeeNotFound struct {
	EmployeeID string
}

func (e *EmployeeNotFound) Error() string {
	return fmt.Sprintf("employee(%s) not found", e.EmployeeID)
}

type EmployeeService interface {
	GetEmployee(ctx context.Context, id string) (*Employee, error)
	ListEmployees(ctx context.Context) ([]Employee, error)
}
