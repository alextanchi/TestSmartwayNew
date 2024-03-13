package repository

import (
	"TestSmartwayNew/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Repository interface {

	//аргументы для методов выбираем в соответствии с ТЗ
	//(например DeleteEmployee удаляет сотрудника по айди => в качестве аргумента будет принимать айди int)

	AddEmployee(employee models.Employee, departmentId string) (string, error)
	DeleteEmployee(id string) error
	ListEmployeeByCompanyId(companyId int) ([]models.Employee, error)
	ListEmployeeByDepartment(departmentName string) ([]models.Employee, error)
	UpdateEmployee(employee models.Employee, departmentId string) error

	GetDepartmentId(phone, name string) (string, error)
	GetEmployeeId(id string) (string, error)
}

type Employee struct {
	db *sql.DB
}

func NewEmployee(db *sql.DB) Repository { //конструктор

	return &Employee{
		db: db,
	}
}

// ConnectDb добавляем подключение к базе
func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "password", "employee")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	log.Println("подключились к БД")

	//defer db.Close()
	return db, nil
}

func (e Employee) AddEmployee(employee models.Employee, departmentId string) (string, error) {

	id := uuid.NewString()

	querySql := `INSERT INTO employees
    (id,
     name, 
     surname,
     phone,
     company_id,
     passport_type,
     passport_number,
     department_id
     ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

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
		d.phone 
		FROM employees e 
		JOIN department d ON e.department_id = d.id
		WHERE e.company_id = $1`

	rows, err := e.db.Query(querySql, companyId)
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
	querySql := `SELECT e.id, 
		e.name,  
		e.surname, 
		e.phone, 
		e.company_id,
		e.passport_type, 
		e.passport_number, 
		d.name, 
		d.phone 
		FROM employees e 
		JOIN department d ON e.department_id = d.id
		WHERE d.name = $1"`

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
	return employees, nil
}

func (e Employee) UpdateEmployee(employee models.Employee, departmentId string) error {

	querySql := "UPDATE employees SET "
	if employee.Name != "" {
		querySql += "name = '" + employee.Name + "',"
	}
	if employee.Surname != "" {
		querySql += "surname = '" + employee.Surname + "',"
	}
	if employee.Phone != "" {
		querySql += "phone = '" + employee.Phone + "',"
	}
	//предобразуем тип
	companyIdString := strconv.Itoa(employee.CompanyId)
	if employee.CompanyId != 0 {
		querySql += "company_id = " + companyIdString + ","
	}
	if employee.Passport.Type != "" {
		querySql += "passport_type = '" + employee.Passport.Type + "',"
	}
	if employee.Passport.Number != "" {
		querySql += "passport_number = '" + employee.Passport.Number + "',"
	}
	if departmentId != "" {
		querySql += "department_id = '" + departmentId + "'"
	}
	querySql += " where id = ?"

	_, err := e.db.Exec(querySql, employee.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Employee) GetDepartmentId(phone, name string) (string, error) {
	querySelect := `SELECT id FROM department WHERE phone = $1 AND name = $2`
	rows, err := e.db.Query(querySelect, phone, name)
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
	querySelect := `SELECT id FROM employees WHERE id = $1`
	rows, err := e.db.Query(querySelect, id)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err == sql.ErrNoRows {
			return "", errors.New("Нет сотрудника с данным id")
		}
		if err != nil {
			return "", err
		}

	}
	return id, nil
}
