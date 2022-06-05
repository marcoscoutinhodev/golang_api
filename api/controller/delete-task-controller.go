package controller

import (
	"encoding/json"
	"golang_api/api/repository"
	"net/http"
	"strconv"
)

func DeleteTaskController(rw http.ResponseWriter, r *http.Request) {
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

	result, err := repository.DeleteTaskByIdRepository(idInteger)

	if err != nil {
		rw.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(rw).Encode("Internal error")
		return
	}

	if count, _ := result.RowsAffected(); count == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode("No task found")
		return
	}

	json.NewEncoder(rw).Encode("Task deleted successfully")
}
