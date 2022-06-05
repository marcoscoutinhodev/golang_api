package controller

import (
	"encoding/json"
	"golang_api/api/model"
	"golang_api/api/repository"
	"net/http"
)

func GetTasksController(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		rw.Header().Set("Allow", "Get")
	}

	rw.Header().Set("Content-Type", "application/json")

	rows, err := repository.GetTasksRepository()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode("Error to get tasks..")
		return
	}

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done)

		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		json.NewEncoder(rw).Encode("No tasks registered")
		return
	}

	json.NewEncoder(rw).Encode(tasks)
}
