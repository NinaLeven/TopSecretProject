package employee

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type employeeClient struct {
	baseURL *url.URL
	client  *http.Client
}

func NewEmployeeServiceClient(client *http.Client, baseURLstring string) (projectmanager.EmployeeService, error) {
	baseURL, err := url.Parse(baseURLstring)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base url: %w", err)
	}
	baseURL.Path = path.Join(baseURL.Path, "api")

	return &employeeClient{
		client:  client,
		baseURL: baseURL,
	}, nil
}

type employee struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Role       string `json:"role"`
}

type employeeList struct {
	Data []employee `json:"data,omitempty"`
}

func (e *employeeClient) GetEmployee(ctx context.Context, id string) (*projectmanager.Employee, error) {
	URL := *e.baseURL
	URL.Path = path.Join(URL.Path, "employees", id)

	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to make request: %w", err)
	}

	res, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("unable to send employees/id request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode == http.StatusNotFound {
		return nil, &projectmanager.EmployeeNotFound{
			EmployeeID: id,
		}
	}
	if res.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("request failed: internal server error")
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: unexpected status: %d", res.StatusCode)
	}

	employee := employee{}
	err = json.NewDecoder(res.Body).Decode(&employee)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal employee: %w", err)
	}

	return employeeToModel(&employee)
}

func (e *employeeClient) ListEmployees(ctx context.Context) ([]projectmanager.Employee, error) {
	URL := *e.baseURL
	URL.Path = path.Join(URL.Path, "employees")

	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to make request: %w", err)
	}

	res, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("unable to send employees request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: unexpected status: %d", res.StatusCode)
	}

	employees := employeeList{}
	err = json.NewDecoder(res.Body).Decode(&employees)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal employees: %w", err)
	}

	return employeesToModel(employees.Data)
}

func employeesToModel(es []employee) ([]projectmanager.Employee, error) {
	res := make([]projectmanager.Employee, 0, len(es))
	for i := range es {
		em, err := employeeToModel(&es[i])
		if err != nil {
			return nil, err
		}
		res = append(res, *em)
	}
	return res, nil
}

func employeeToModel(e *employee) (*projectmanager.Employee, error) {
	if e == nil {
		return nil, nil
	}

	role, err := projectmanager.EmployeeRoleFromString(e.Role)
	if err != nil {
		return nil, err
	}

	return &projectmanager.Employee{
		Id:         e.Id,
		FirstName:  e.FirstName,
		LastName:   e.LastName,
		Email:      e.Email,
		Department: e.Department,
		Role:       role,
	}, nil
}
