package controller

import (
	"encoding/json"
	"golang_api/api/model"
	"golang_api/api/repository"
	"net/http"
)

func AddTaskController(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "Post")
	}

	bodyDecoder := json.NewDecoder(r.Body)

	var taskModel model.Task
	err := bodyDecoder.Decode(&taskModel)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode("Internal error")
		return
	}

	err = taskModel.Validate()

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	err = repository.AddTaskRepository(taskModel)

	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(rw).Encode("Internal error")
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
