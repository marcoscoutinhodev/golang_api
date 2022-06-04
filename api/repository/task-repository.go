package repository

import (
	database "golang_api/api/database/sql"
	"golang_api/api/model"
)

func AddTaskRepository(task model.Task) error {
	db, err := database.Connection()

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO golang_db.tbTask(title, description, done) VALUES(?, ?, ?)")
	stmt.Exec(task.Title, task.Description, false)

	return nil
}