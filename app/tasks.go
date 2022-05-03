package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)

type task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Task One",
		Content: "Lorem Ipsum",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to the API")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(
			Message("Error, insert a valid task"),
		)
		return
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.WriteHeader((http.StatusCreated))
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(
			Message("Error, invalid ID"),
		)
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	json.NewEncoder(w).Encode(
		Message("Task not found"),
	)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode(
			Message("Error, invalid ID"),
		)
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			json.NewEncoder(w).Encode(
				Message("The task has been removed succesfully"),
			)
			return
		}
	}

	json.NewEncoder(w).Encode(
		Message("Task not found"),
	)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	var updatedTask task

	if err != nil {
		json.NewEncoder(w).Encode(
			Message("Error, invalid ID"),
		)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(
			Message("Please Enter Valid Data"),
		)
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			json.NewEncoder(w).Encode(
				Message("The task has been updated succesfully"),
			)
			return
		}
	}

	json.NewEncoder(w).Encode(
		Message("Task not found"),
	)
}