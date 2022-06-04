package model

import "errors"

type Task struct {
	Id          *int   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var (
	ErrorTaskTitleEmpty       = errors.New("Task.Title can't be empty")
	ErrorTaskDescriptionEmpty = errors.New("Task.Description can't be empty")
)

func (t Task) Validate() error {
	if t.Title == "" {
		return ErrorTaskTitleEmpty
	}

	if t.Description == "" {
		return ErrorTaskDescriptionEmpty
	}

	return nil
}
