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
type ResponseCreated struct {
	id string `json:"id"`
}
type ResponseError struct {
	statusCode   int    `json:"statusCode"`
	errorMessage string `json:"errorMessage"`
}

func NewController(srv service.Service) Controller {
	return &EmployeeController{
		useCase: srv,
	}
}

func (c EmployeeController) AddEmployee(ctx *gin.Context) {
	var id string
	employee := models.Employee{}
	err := ctx.ShouldBind(&employee)
	if err != nil {
		log.Println(err)
		//	returnError(ctx, http.StatusBadRequest, "Неверный формат входных данных")
		ctx.String(http.StatusBadRequest, "Неверный формат входных данных")
		return
	}
	log.Println("controller-AddEmployee-checkpoint-1")

	id, err = c.useCase.AddEmployee(employee)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "Ошибка при добавлении сотрудника")
		return
	}
	log.Println("controller-AddEmployee-checkpoint-2")
	ctx.String(http.StatusCreated, "Добавлен сотрудник id: "+id)

}

func (c EmployeeController) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.useCase.DeleteEmployee(id)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "Ошибка при удалении сотрудника") //или здесь 400?
		return
	}
	log.Println("controller-DeleteEmployee-checkpoint-1")
	ctx.JSON(http.StatusNoContent, "Сотрудник удален")

}

func (c EmployeeController) ListEmployeeByCompanyId(ctx *gin.Context) {
	companyId := ctx.Param("companyid")
	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, "Ошибка при вводе companyid")
		return
	}
	log.Println("controller-ListEmployeeByCompanyId-checkpoint-1")
	result, err := c.useCase.ListEmployeeByCompanyId(companyIdInt)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "Ошибка при выводе списка сотрудников для указанной компании ")
		return
	}
	log.Println("controller-ListEmployeeByCompanyId-checkpoint-2")
	ctx.JSON(http.StatusOK, result)

}

func (c EmployeeController) ListEmployeeByDepartment(ctx *gin.Context) {
	department := ctx.Param("department")
	log.Println("controller-ListEmployeeByDepartment-checkpoint-0")
	_, err := c.useCase.ListEmployeeByDepartment(department)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "Ошибка при выводе списка сотрудников для указанного отдела ")
		return
	}
	log.Println("controller-ListEmployeeByDepartment-checkpoint-1")
	ctx.String(http.StatusOK, "Список сотрудников для указанного отдела компании")

}

func (c EmployeeController) UpdateEmployee(ctx *gin.Context) {
	employee := models.Employee{}
	err := ctx.ShouldBind(employee)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("controller-UpdateEmployee-checkpoint-1")
	err = c.useCase.UpdateEmployee(employee)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка при изменении сотрудника ")
		return
	}
	log.Println("controller-UpdateEmployee-checkpoint-2")
	ctx.JSON(http.StatusOK, "Данные пользователя успешно обновлены")

}

func ReturnError(ctx gin.Context, statusCode int, errorMessage string) {
	responseError := &ResponseError{
		statusCode:   statusCode,
		errorMessage: errorMessage,
	}
	ctx.JSON(statusCode, responseError)
}
