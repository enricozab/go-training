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
	err := http.ListenAndServe(":8080", db.handler())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int

		// Checks if URL path is correct. If correct, leads to specific db process, else it will log an error
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		} else {
			http.Error(w, "Incorrect URL", http.StatusBadRequest)
		}
	}
}

func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Set("Content-Type", "application/json")
		var contact Contact

		// Decodes the request body
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Adds new contact to the database
		db.mu.Lock()

		for _, item := range db.contacts {
			if item.First == contact.First && item.Last == contact.Last {
				defer db.mu.Unlock()

				http.Error(w, "Contact already exists", http.StatusConflict)

				return
			}
		}

		contact.ID = db.nextID
		db.nextID++
		db.contacts = append(db.contacts, contact)
		db.mu.Unlock()

		w.WriteHeader(http.StatusCreated)

		fmt.Fprintln(w, "Success: New contact has been added successfully.")
		data, _ := json.Marshal(contact)
		fmt.Fprintf(w, "%+v\n", string(data))
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintln(w, "Contacts")

		// Returns all contacts
		if err := json.NewEncoder(w).Encode(db.contacts); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		// Other methods are not allowed
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		// Returns a specific contact if provided id exists
		db.mu.Lock()
		for _, item := range db.contacts {
			if id == item.ID {
				fmt.Fprintf(w, "Contact ID %v\n", id)

				defer db.mu.Unlock()

				if err := json.NewEncoder(w).Encode(item); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				return
			}
		}
		db.mu.Unlock()

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case "DELETE":
		w.Header().Set("Content-Type", "application/json")

		// Deletes a contact from the database if provided id exists
		db.mu.Lock()
		for j, item := range db.contacts {
			if id == item.ID {
				db.contacts = append(db.contacts[:j], db.contacts[j+1:]...)

				defer db.mu.Unlock()

				fmt.Fprintf(w, "Success: Contact with ID %d has been deleted successfully.", id)

				return
			}
		}
		db.mu.Unlock()

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	case "PUT":
		w.Header().Set("Content-Type", "application/json")

		var contact Contact

		// Decodes the request body
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Updates a contact from the database if provided id exists
		db.mu.Lock()
		for i, item := range db.contacts {
			if id == item.ID {
				contact.ID = item.ID
				db.contacts[i] = contact

				defer db.mu.Unlock()

				fmt.Fprintf(w, "Success: Contact with ID %d has been updated successfully.", id)

				return
			}
		}
		db.mu.Unlock()

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	default:
		// Other methods are not allowed
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
