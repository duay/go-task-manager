package main

import (
	"database/sql"
	"log"
	"time"
)

func StartWorker(db *sql.DB) {
	for {
		time.Sleep(1 * time.Minute)
		_, err := db.Exec("UPDATE tasks SET archived = 1 WHERE done = 1")
		if err != nil {
			log.Println("Worker error:", err)
		} else {
			log.Println("Archived done tasks.")
		}
	}
}
