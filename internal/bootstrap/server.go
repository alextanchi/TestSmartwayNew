package bootstrap

import (
	"TestSmartwayNew/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.Controller
}

// NewServer объединили  пакет http и контроллер
func NewServer(cnt controller.Controller) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20,          //1MB
			ReadTimeout:    10 * time.Second, //10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt: cnt,
	}
}

// InitRoutes инициализируем все наши эндпоинты
func (s Server) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/employee", s.cnt.AddEmployee)
	router.DELETE("/employee/:id", s.cnt.DeleteEmployee)
	router.GET("/employee/list/:companyid", s.cnt.ListEmployeeByCompanyId)
	router.GET("/employee/list/:department", s.cnt.ListEmployeeByDepartment)
	router.PUT("/employee", s.cnt.UpdateEmployee)

	return router

}


func (s *Server) Run(port string, handler http.Handler) error { 

	return s.httpServer.ListenAndServe() 

func (s *Server) Shutdown(ctx context.Context) error { // остановка сервера
	return s.httpServer.Shutdown(ctx)
}*/
