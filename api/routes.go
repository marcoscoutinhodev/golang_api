package api

import (
	"golang_api/api/controller"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/task/add", controller.AddTaskController)
	mux.HandleFunc("/task/get", controller.GetTaskController)
	mux.HandleFunc("/task/update", controller.UpdateTaskController)
	mux.HandleFunc("/task/delete", controller.DeleteTaskController)
	mux.HandleFunc("/task/toggle-done-field", controller.ToggleDoneFieldTaskController)

	return mux
}
