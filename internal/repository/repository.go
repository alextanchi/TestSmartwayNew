package repository

import (
	"TestSmartwayNew/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type Repository interface {

	//аргументы для методов выбираем в соответствии с ТЗ
	//(например DeleteEmployee удаляет сотрудника по айди => в качестве аргумента будет принимать айди int)

	AddEmployee(employee models.Employee, departmentId string) (string, error) //принимаем всю структуру employee в качестве аргумента
	DeleteEmployee(id string) error
	ListEmployeeByCompanyId(companyId int) ([]models.Employee, error)
	ListEmployeeByDepartment(departmentName string) ([]models.Employee, error)
	UpdateEmployee(employee models.Employee, departmentId string) error
	GetDepartmentId(phone, name string) (string, error)
	GetEmployeeId(id string) (string, error)
}

//создаем конструктор и добавляем в нем подключение к базе

type Employee struct {
	db *sql.DB
}

func NewEmployee(db *sql.DB) Repository { //конструктор

	return &Employee{
		db: db,
	}
}

func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 54321, "postgres", "postgres1234", "employees")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	return db, nil
}

func (e Employee) AddEmployee(employee models.Employee, departmentId string) (string, error) {

	id := uuid.NewString()
	querySql := `INSERT INTO empolyees
    (id,
     name, 
     surname,
     phone,
     company_id,
     passport_type,
     passport_number,
     department_id
     ) VALUES ($1, $2, $3, $4, $5, $6, $7, &8)`

	_, err := e.db.Exec(querySql,
		id,
		employee.Name,
		employee.Surname,
		employee.Phone,
		employee.CompanyId,
		employee.Passport.Type,
		employee.Passport.Number,
		departmentId,
	)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (e Employee) DeleteEmployee(id string) error {
	_, err := e.db.Exec("DELETE FROM employees WHERE id = $1", id)
	return err
}

func (e Employee) ListEmployeeByCompanyId(companyId int) ([]models.Employee, error) {
	querySql := `SELECT 
    	e.id,
		e.name, 
		e.surname, 
		e.phone, 
		e.company_id, 
		e.passport_type, 
		e.passport_number,
		d.name, 
		d.phone, 
		FROM employees e, 
		JOIN department d ON e.id = d.employee_id, 
		WHERE e.company_id = $1`

	rows, err := e.db.Query(querySql, companyId) // вопрос
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]models.Employee, 0)
	for rows.Next() {
		emp := models.Employee{}
		err = rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Surname,
			&emp.Phone,
			&emp.CompanyId,
			&emp.Passport.Type,
			&emp.Passport.Number,
			&emp.Department.Name,
			&emp.Department.Phone,
		)
		if err == sql.ErrNoRows {
			return []models.Employee{}, nil
		}
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (e Employee) ListEmployeeByDepartment(departmentName string) ([]models.Employee, error) {
	querySql := "SELECT e.id, \n" +
		"e.name, \n" +
		"e.surname, \n" +
		"e.phone, \n" +
		"e.company_id, \n" +
		"e.passport_type, \n" +
		"e.passport_number, \n" +
		"d.name, \n" +
		"d.phone, \n" +
		"FROM employees e \n" +
		"JOIN department d ON e.id = d.employee_id, \n" +
		"WHERE d.name = $1"

	rows, err := e.db.Query(querySql, departmentName) // вопрос
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []models.Employee{}
	for rows.Next() {
		emp := models.Employee{}
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Surname,
			&emp.Phone,
			&emp.CompanyId,
			&emp.Passport.Type,
			&emp.Passport.Number,
			&emp.Department.Name,
			&emp.Department.Phone,
		)
		if err == sql.ErrNoRows {
			return []models.Employee{}, nil
		}
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return []models.Employee{}, nil
}

func (e Employee) UpdateEmployee(employee models.Employee, departmentId string) error {

	querySql := "UPDATE employee SET "
	if employee.Name != "" {
		querySql += "name = " + employee.Name //нужно ведь запятые после каждого поля добавить?
	}
	if employee.Surname != "" {
		querySql += "surname = " + employee.Surname
	}
	if employee.Phone != "" {
		querySql += "phone = " + employee.Phone
	}
	companyIdString := strconv.Itoa(employee.CompanyId) //предобразуем тип
	if employee.CompanyId != 0 {
		querySql += "company_id = " + companyIdString
	}
	if employee.Passport.Type != "" {
		querySql += "passport_type = " + employee.Passport.Type
	}
	if employee.Passport.Number != "" {
		querySql += "passport_number = " + employee.Passport.Number
	}
	if departmentId != "" {
		querySql += "department_id = " + departmentId
	}
	querySql += " where id = ?"
}

func (e Employee) GetDepartmentId(phone, name string) (string, error) {
	querySelect := `SELECT id FROM department WHERE phone = ? AND name = ?`
	rows, err := e.db.Query(querySelect, phone, name) // вопрос
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var departmentId string
	for rows.Next() {
		err := rows.Scan(&departmentId)
		if err == sql.ErrNoRows {
			return "", errors.New("Не существующий департамент")
		}
		if err != nil {
			return "", err
		}

	}
	return departmentId, nil
}
func (e Employee) GetEmployeeId(id string) (string, error) {
	/*querySelect := `SELECT id FROM department WHERE phone = ? AND name = ?`
		rows, err := e.db.Query(querySelect, phone, name) // вопрос
		if err != nil {
			return "", err
		}
		defer rows.Close()

		var departmenId string
		for rows.Next() {
			err := rows.Scan(&departmenId)
			if err == sql.ErrNoRows {
				return "", errors.New("Не существующий департамент")
			}
			if err != nil {
				return "", err
			}

		}
		return departmenId, nil
	}*/
}
