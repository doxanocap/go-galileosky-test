package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo/internal/model"
)

var ID int64

func TestCreateTodoItem(t *testing.T) {
	router := getRouter()

	todoItem := model.TodoItem{
		Title:       "Sample TodoItem",
		Description: "This is a sample TodoItem",
		Done:        false,
		CreatedAt:   &time.Time{},
		DeletedAt:   &time.Time{},
	}
	payload, _ := json.Marshal(todoItem)

	r, err := http.NewRequest("POST", "/v1/todo", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	res := &model.TodoItem{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	ID = res.Id
	assert.NotEmpty(t, res)
}

func TestGetAllTodoItems(t *testing.T) {
	router := getRouter()

	r, err := http.NewRequest("GET", "/v1/todo/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	res := &[]model.TodoItem{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, res)
}

func TestGetByIDTodoItems(t *testing.T) {
	router := getRouter()

	r, err := http.NewRequest("GET", fmt.Sprintf("/v1/todo/%d", ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body)

	res := &model.TodoItem{}
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, res)
}

func TestDeleteByIDTodoItems(t *testing.T) {
	router := getRouter()

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/v1/todo/%d", ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
