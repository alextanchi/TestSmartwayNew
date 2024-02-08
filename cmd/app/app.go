package app

import "TestSmartwayNew/internal/repository"

func Run() error {
	db, err := connectDb()
	if err != nil {
		return err
	}
	repository.NewEmployee(db)
	return nil
}
