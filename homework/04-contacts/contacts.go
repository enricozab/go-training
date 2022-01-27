package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Contact struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

type Database struct {
	nextID   int
	mu       sync.Mutex
	contacts []Contact
}

func main() {
	db := &Database{contacts: []Contact{}}
	db.nextID = 1
	http.ListenAndServe(":8080", db.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int

		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		} else {
			log.Fatalln("Incorrect URL.")
		}
	}
}

func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var contact Contact

		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.mu.Lock()
		contact.ID = db.nextID
		db.nextID++
		db.contacts = append(db.contacts, contact)
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintln(w, "Success: New contact has been added successfully.")
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintln(w, "Contacts")

		if err := json.NewEncoder(w).Encode(db.contacts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		isExisting := false

		db.mu.Lock()
		for j, item := range db.contacts {
			if id == item.ID {
				isExisting = true
				db.contacts = append(db.contacts[:j], db.contacts[j+1:]...)

				break
			}
		}
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		if isExisting {
			fmt.Fprintf(w, "Success: Contact with ID %d has been deleted successfully.", id)
		} else {
			fmt.Fprintf(w, "Failed: Contact with ID %d does not exist.", id)
		}
	case "PUT":
		var contact Contact

		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contactIndex := -1

		db.mu.Lock()
		for i, item := range db.contacts {
			if id == item.ID {
				contactIndex = i
				contact.ID = item.ID

				break
			}
		}
		if contactIndex >= 0 {
			db.contacts[contactIndex] = contact
		}
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		if contactIndex >= 0 {
			fmt.Fprintf(w, "Success: Contact with ID %d has been updated successfully.", id)
		} else {
			fmt.Fprintf(w, "Failed: Contact with ID %d does not exist.", id)
		}
	}
}
