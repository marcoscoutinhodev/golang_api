package repository

import (
	"database/sql"
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

func GetTasksRepository() (*sql.Rows, error) {
	db, err := database.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM golang_db.tbTask t ORDER BY t.title ASC")

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetTaskByIdRepository(id int) (*sql.Row, error) {
	db, err := database.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT * FROM golang_db.tbTask t WHERE t.id = ?", id)

	return row, nil
}
