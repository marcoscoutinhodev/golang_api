package controller

import (
	"encoding/json"
	"golang_api/api/model"
	"golang_api/api/repository"
	"net/http"
	"strconv"
)

func ToggleDoneFieldTaskController(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "Post")
	}

	rw.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	if id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("Task ID must be provided")
		return
	}

	idInteger, err := strconv.Atoi(id)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("Task ID invalid")
		return
	}

	row, err := repository.GetTaskByIdRepository(idInteger)

	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(rw).Encode("Internal error")
		return
	}

	var taskModel model.Task
	err = row.Scan(&taskModel.Id, &taskModel.Title, &taskModel.Description, &taskModel.Done)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("No task found")
		return
	}

	taskModel.Done = !taskModel.Done

	err = repository.UpdateTaskRepository(taskModel)

	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(rw).Encode("Internal error to update task")
		return
	}

	json.NewEncoder(rw).Encode("Task updated successfully")
}
