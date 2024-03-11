package controller

import (
	"TestSmartwayNew/internal/models"
	"TestSmartwayNew/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Controller interface {
	AddEmployee(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
	ListEmployeeByCompanyId(ctx *gin.Context)
	ListEmployeeByDepartment(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
}
type EmployeeController struct {
	useCase service.Service
}

func NewController(srv service.Service) Controller {
	return &EmployeeController{
		useCase: srv,
	}
}

func (c EmployeeController) AddEmployee(ctx *gin.Context) {
	var id string
	employee := models.Employee{}
	err := ctx.ShouldBind(employee)
	if err != nil {
		return
	}
	log.Println("controller-AddEmployee-checkpoint-1")

	//не уверен что правильно обработал ошибку (типо может нужно одну обработку в конце сделать?)
	id, err = c.useCase.AddEmployee(employee)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при добавлении сотрудника")
		return
	}
	ctx.String(http.StatusCreated, "Добавлен сотрудник id: "+id)
	log.Println("controller-AddEmployee-checkpoint-2")
}

func (c EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.useCase.DeleteEmployee(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при удалении сотрудника") //или здесь 400?
		return
	}
	ctx.JSON(http.StatusNoContent, "Сотрудник удален")
	log.Println("controller-DeleteEmployee-checkpoint-1")
}

func (c EmployeeController) ListEmployeeByCompanyId(ctx *gin.Context) {
	companyId := ctx.Param("companyid")
	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		return
	}
	log.Println("controller-ListEmployeeByCompanyId-checkpoint-1")
	_, err = c.useCase.ListEmployeeByCompanyId(companyIdInt)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при выводе списка сотрудников для указанной компании ")
		return
	}
	ctx.String(http.StatusOK, "Список сотрудников для указанной компании")
	log.Println("controller-ListEmployeeByCompanyId-checkpoint-2")
}

func (c EmployeeController) ListEmployeeByDepartment(ctx *gin.Context) {
	department := ctx.Param("department")
	_, err := c.useCase.ListEmployeeByDepartment(department)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при выводе списка сотрудников для указанного отдела ")
		return
	}
	ctx.String(http.StatusOK, "Список сотрудников для указанного отдела компании")
	log.Println("controller-ListEmployeeByDepartment-checkpoint-1")
}

func (c EmployeeController) UpdateEmployee(ctx *gin.Context) {
	employee := models.Employee{}
	err := ctx.ShouldBind(employee)
	if err != nil {
		return
	}
	log.Println("controller-UpdateEmployee-checkpoint-1")
	err = c.useCase.UpdateEmployee(employee)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при изменении сотрудника ")
		return
	}
	ctx.JSON(http.StatusOK, "Данные пользователя успешно обновлены")
	log.Println("controller-UpdateEmployee-checkpoint-2")
}
