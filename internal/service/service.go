package service

import (
	"TestSmartwayNew/internal/models"
	"TestSmartwayNew/internal/repository"
)

type Service interface {
	AddEmployee(employee models.Employee) (string, error)
	DeleteEmployee(id string) error
	ListEmployeeByCompanyId(companyId int) ([]models.Employee, error)
	ListEmployeeByDepartment(departmentName string) ([]models.Employee, error)
	UpdateEmployee(employee models.Employee) error
}
type EmployeeService struct {
	storage repository.Repository
}

func NewService(repos repository.Repository) Service {

	return EmployeeService{
		storage: repos,
	}
}

func (s EmployeeService) AddEmployee(employee models.Employee) (string, error) {
	//подразумеваем, что department-ы уже создан и проверяем существует ли department указанный в запросе
	departmentId, err := s.storage.GetDepartmentId(employee.Department.Phone, employee.Department.Name)
	if err != nil {
		return "", err
	}
	return s.storage.AddEmployee(employee, departmentId)

}
func (s EmployeeService) DeleteEmployee(id string) error {
	_, err := s.storage.GetEmployeeId(id) //проверка на существующий id
	if err != nil {
		return err
	}

	return s.storage.DeleteEmployee(id)

}
func (s EmployeeService) ListEmployeeByCompanyId(companyId int) ([]models.Employee, error) {

	return s.storage.ListEmployeeByCompanyId(companyId)

}

func (s EmployeeService) ListEmployeeByDepartment(departmentName string) ([]models.Employee, error) {
	return s.storage.ListEmployeeByDepartment(departmentName)
}
func (s EmployeeService) UpdateEmployee(employee models.Employee) error {
	var departmentId string
	var err error
	if employee.Phone != "" && employee.Name != "" {
		departmentId, err = s.storage.GetDepartmentId(employee.Phone, employee.Name)
		if err != nil {
			return err
		}
	}

	return s.storage.UpdateEmployee(employee, departmentId)
}
