package main

import (
	"log"
	"net/http"
)

func main() {
	// Init DB
	db := InitDB()
	defer db.Close()

	// Run background worker
	go StartWorker(db)

	// Routes
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetTasksHandler(w, r, db)
		} else if r.Method == http.MethodPost {
			CreateTaskHandler(w, r, db)
		}
	})

	http.HandleFunc("/tasks/done", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			MarkDoneHandler(w, r, db)
		}
	})

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
