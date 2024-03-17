package app

import (
	"TestSmartwayNew/internal/bootstrap"
	"TestSmartwayNew/internal/controller"
	"TestSmartwayNew/internal/repository"
	"TestSmartwayNew/internal/service"
	"log"
)

func Run() error {

	db, err := repository.ConnectDb()
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("app.go-Run-Чекпоинт 1")
	store := repository.NewEmployee(db)

	srv := service.NewService(store)

	cnt := controller.NewController(srv)

	serv := bootstrap.NewServer(cnt)

	router := serv.InitRoutes()

	router.Run(":8080")

	return nil
}
