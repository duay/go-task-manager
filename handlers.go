package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, title, done, archived FROM tasks WHERE archived = 0")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.Done, &t.Archived)
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var t Task
	json.NewDecoder(r.Body).Decode(&t)

	_, err := db.Exec("INSERT INTO tasks (title) VALUES (?)", t.Title)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func MarkDoneHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var t Task
	json.NewDecoder(r.Body).Decode(&t)

	_, err := db.Exec("UPDATE tasks SET done = 1 WHERE id = ?", t.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
