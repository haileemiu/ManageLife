package model

import "time"

type TaskItemResponse struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Notes          string    `json:"notes"`
	IsTimeSenstive bool      `json:"isTimeSensitve"`
	IsImportant    bool      `json:"isImportant"`
	RemindAt       time.Time `json:"remindAt"`
	DueAt          time.Time `json:"dueAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type TaskListResponse struct {
	Tasks []TaskItemResponse `json:"tasks"`
}

type TaskCreateRequest struct {
	Title          string    `json:"title"`
	Notes          string    `json:"notes"`
	IsTimeSenstive bool      `json:"isTimeSensitve"`
	IsImportant    bool      `json:"isImportant"`
	RemindAt       time.Time `json:"remindAt"`
	DueAt          time.Time `json:"dueAt"`
}

type TaskUpdateRequest struct {
	Title          string    `json:"title"`
	Notes          string    `json:"notes"`
	IsTimeSenstive bool      `json:"isTimeSensitve"`
	IsImportant    bool      `json:"isImportant"`
	RemindAt       time.Time `json:"remindAt"`
	DueAt          time.Time `json:"dueAt"`
}

func (tcr TaskCreateRequest) Validate() (bool, map[string][]string) {
	errs := map[string][]string{}

	if tcr.Title == "" {
		errs["title"] = append(errs["title"], "title must be set")
	}

	if len(tcr.Title) < 3 || len(tcr.Title) > 120 {
		errs["title"] = append(errs["title"], "title must be between 3 and 120 characters")
	}

	// TODO: put other validations here

	return len(errs) == 0, errs
}
