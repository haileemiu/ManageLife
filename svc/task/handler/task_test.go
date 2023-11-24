package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"

	"github.com/haileemiu/manage-life/ent"
	"github.com/haileemiu/manage-life/ent/enttest"
	"github.com/haileemiu/manage-life/svc/task/handler"
	"github.com/haileemiu/manage-life/svc/task/model"
)

func createHandler(t *testing.T) (http.Handler, *ent.Client) {
	testEntClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")

	r := chi.NewRouter()
	hdl := handler.New(testEntClient)

	hdl.Routes(r)

	return r, testEntClient
}

func TestGetByID(t *testing.T) {
	// Arrange
	r, entClient := createHandler(t)
	defer entClient.Close()

	taskItemRes := &model.TaskItemResponse{
		Title: "task_title",
		Notes: "task_notes",
	}

	task, err := entClient.Task.Create().
		SetTitle(taskItemRes.Title).
		SetNotes(taskItemRes.Notes).
		Save(context.Background())
	assert.NoError(t, err)

	// Testing Table
	tests := []struct {
		name string
		id   string

		shouldCloseEnt bool

		expectedStatus int
		expectedBody   *model.TaskItemResponse
	}{
		{
			name:           "success",
			id:             fmt.Sprint(task.ID),
			expectedStatus: http.StatusOK,
			expectedBody:   taskItemRes,
		},
		{
			name:           "not found",
			id:             fmt.Sprint(task.ID + 1),
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "bad request",
			id:             "sweet",
			expectedStatus: http.StatusBadRequest,
		},
		// db closed tests
		{
			name:           "ent error",
			id:             fmt.Sprint(task.ID),
			shouldCloseEnt: true,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		// Act
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCloseEnt {
				entClient.Close()
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("/%s", tt.id), nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			// Assert
			if tt.expectedBody != nil {
				resItem := model.TaskItemResponse{}
				err = json.NewDecoder(rr.Body).Decode(&resItem)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody.Title, resItem.Title)
				assert.Equal(t, tt.expectedBody.Notes, resItem.Notes)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	// Arrange
	r, entClient := createHandler(t)
	defer entClient.Close()

	type RequestBody struct {
		Title string `json:"title"`
		Notes string `json:"notes"`
	}

	// Testing Table
	tests := []struct {
		name string
		body RequestBody

		shouldCloseEnt bool

		expectedStatus int
		expectedBody   *model.TaskItemResponse
	}{
		{
			name: "Success Test",
			body: RequestBody{Title: "test_title", Notes: "test_notes"},

			expectedStatus: http.StatusCreated,
			expectedBody: &model.TaskItemResponse{
				Title: "test_title",
				Notes: "test_notes",
			},
		},
		// {
		// 	name:           "not found",
		// 	id:             fmt.Sprint(task.ID + 1),
		// 	expectedStatus: http.StatusNotFound,
		// },
		// {
		// 	name:           "bad request",
		// 	id:             "sweet",
		// 	expectedStatus: http.StatusBadRequest,
		// },
		// // db closed tests
		// {
		// 	name:           "ent error",
		// 	id:             fmt.Sprint(task.ID),
		// 	shouldCloseEnt: true,
		// 	expectedStatus: http.StatusInternalServerError,
		// },
	}

	for _, tt := range tests {
		// Act
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldCloseEnt {
				entClient.Close()
			}

			reqJson, err := json.Marshal(tt.body)
			assert.NoError(t, err)

			taskReader := bytes.NewReader(reqJson)

			req, err := http.NewRequest("POST", "/", taskReader)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if tt.expectedBody != nil {
				resItem := model.TaskItemResponse{}
				err = json.NewDecoder(rr.Body).Decode(&resItem)
				assert.NoError(t, err)

				assert.NotNil(t, tt.expectedBody.ID)
				assert.Equal(t, tt.expectedBody.Title, resItem.Title)
				assert.Equal(t, tt.expectedBody.Notes, resItem.Notes)
			}
		})
	}
}
