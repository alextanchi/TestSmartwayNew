package bootstrap

import (
	"TestSmartwayNew/internal/controller"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.Controller
}

// NewServer объединили  пакет http и контроллера
func NewServer(cnt controller.Controller) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,          //1MB
			ReadTimeout:    10 * time.Second, //10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt: cnt,
	}
}

/*
func (s *Server) Run(port string, handler http.Handler) error { //запуск, добавили в метод run аргумент хендлер типа интерфейса хендлер

	return s.httpServer.ListenAndServe() //слушаем все входящие запросы для дальнейшей обработки
}
func (s *Server) Shutdown(ctx context.Context) error { // остановка сервера
	return s.httpServer.Shutdown(ctx)
}*/
