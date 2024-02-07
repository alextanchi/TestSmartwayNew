package repository

import (
	_ "TestSmartwayNew/internal/config"
	"TestSmartwayNew/internal/models"
	"database/sql"
	"fmt"
)

type Repository interface {

	//аргументы для методов выбираем в соответствии с ТЗ
	//(например DeleteEmployee удаляет сотрудника по айди => в качестве аргумента будет принимать айди int)

	AddEmployee(employee models.Employee) (int, error) //принимаем всю структуру employee в качестве аргумента
	DeleteEmployee(id int) error
	ListEmployeeByCompanyId(companyId string) ([]models.Employee, error)
	ListEmployeeByDepartment(departmentName string) ([]models.Employee, error)
}

//создаем конструктор и добавляем в нем подключение к базе

type Employee struct {
	db *sql.DB
}

func NewEmployee() Repository { //конструктор

	return &Employee{
		db: db,
	}
}

func connectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	return db, nil
}

func (e Employee) AddEmployee(employee models.Employee) (int, error) {

}

func (e Employee) DeleteEmployee(id int) error {
	result, err := e.db.Exec("delete from employees where id = $1", id)
	if err != nil {

	}
	return err
}

func (e Employee) ListEmployeeByCompanyId(companyId string) ([]models.Employee, error) {

}
func (e Employee) ListEmployeeByDepartment(departmentName string) ([]models.Employee, error) {

}
