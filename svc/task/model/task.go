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
