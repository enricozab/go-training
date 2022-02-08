// Package main implements CRUD for movie list of a user
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Structure of a User
type User struct {
	Id        int
	Username  string
	Password  string
	DateAdded string
}

// Structure of a Movie linked to a User
type Movie struct {
	Id           int
	UserId       int
	Title        string
	IsWatched    bool
	DateAdded    string
	LastModified string
}

// Structure of a Log linked to a User
type Log struct {
	Id          int
	UserId      int
	Description string
	Date        string
}

// Structure of the Database containing the list of users, movies, and logs
type Database struct {
	nextUserId  int
	nextMovieId int
	nextLogId   int
	mu          sync.Mutex
	users       []User
	movieList   []Movie
	logs        []Log
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	// Creates a file (.txt) for logs, or append to the file if existing
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// Creates custom loggers
	InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Initializes the database
	db := &Database{users: []User{}, movieList: []Movie{}, logs: []Log{}}
	db.nextUserId = 1
	db.nextMovieId = 1
	db.nextLogId = 1
	http.ListenAndServe(":8080", db.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		var id int

		// API routes
		if r.URL.Path == "/users" {
			db.processUsers(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/movie-list/%d", &id); n == 1 {
			db.processMovieList(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/movie-list/edit/%d", &id); n == 1 {
			db.processMovieList(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/movie-list/delete/%d", &id); n == 1 {
			db.processMovieList(id, w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/logs/%d", &id); n == 1 {
			db.processLogs(id, w, r)
		} else {
			ErrorLogger.Printf("Incorrect URL: %v\n", r.URL)
		}
	}
}

// enableCORS allows cross-origin requests
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (db *Database) processUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user User
		var tempLog Log

		// Returns error if request body does not fit with the user structure
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checks if username already exists
		var isExisting = false
		for _, u := range db.users {
			if strings.EqualFold(strings.ToLower(u.Username), strings.ToLower(user.Username)) {
				isExisting = true
				break
			}
		}
		if isExisting {
			WarningLogger.Printf("Error: Username %v already exists\n", user.Username)
			http.Error(w, "Error: Username already exists. Please try another one.", http.StatusBadRequest)
			return
		}

		// Adds new user and a log to the database
		db.mu.Lock()

		user.Id = db.nextUserId
		user.DateAdded = time.Now().Format("2006-01-02 15:04:05")
		db.nextUserId++
		db.users = append(db.users, user)

		InfoLogger.Printf("User ID %v has been created\n", user.Id)

		tempLog.Id = db.nextLogId
		tempLog.UserId = user.Id
		tempLog.Description = "Account was created"
		tempLog.Date = time.Now().Format("2006-01-02 15:04:05")
		db.nextLogId++
		db.logs = append(db.logs, tempLog)

		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of users if not error
		if err := json.NewEncoder(w).Encode(db.users); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		InfoLogger.Println("Users have been retrieved")

		// Returns updated list of users if not error
		if err := json.NewEncoder(w).Encode(db.users); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (db *Database) processMovieList(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var movie Movie
		var tempLog Log

		// Returns error if request body does not fit with the movie structure
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checks if movie title already exists in a user
		var isExisting = false
		for _, m := range db.movieList {
			if m.UserId == id && strings.EqualFold(strings.ToLower(m.Title), strings.ToLower(movie.Title)) {
				isExisting = true
				break
			}
		}
		if isExisting {
			WarningLogger.Printf("Error: Movie title %v already exists\n", movie.Title)
			http.Error(w, "Error: Movie title already exists. Please try another one.", http.StatusBadRequest)
			return
		}

		// Adds new movie to a user and a log to the database
		db.mu.Lock()

		movie.Id = db.nextMovieId
		movie.UserId = id
		movie.DateAdded = time.Now().Format("2006-01-02 15:04:05")
		movie.LastModified = time.Now().Format("2006-01-02 15:04:05")
		db.nextMovieId++
		db.movieList = append(db.movieList, movie)

		InfoLogger.Printf("Movie ID %v has been created\n", movie.Id)

		tempLog.Id = db.nextLogId
		tempLog.UserId = id
		tempLog.Description = "Movie " + movie.Title + " was added in the movie list"
		tempLog.Date = time.Now().Format("2006-01-02 15:04:05")
		db.nextLogId++
		db.logs = append(db.logs, tempLog)

		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of movies of a user if not error
		if err := json.NewEncoder(w).Encode(filterMoviesByUser(db, id)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		InfoLogger.Printf("Movies have been retrieved for User ID %v\n", id)

		// Returns updated list of movies of a user if not error
		if err := json.NewEncoder(w).Encode(filterMoviesByUser(db, id)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		var userId int
		var movie Movie
		var tempLog Log

		// Returns error if request body does not fit with the movie structure
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		movieIndex := -1
		var isExisting = false

		// Edit the movie's detailts and adds a log to the database
		db.mu.Lock()
		for i, item := range db.movieList {
			if id == item.Id {
				userId = item.UserId

				// Breaks loop and method if movie title already exists upon updating a movie title
				for _, m := range db.movieList {
					if m.UserId == userId && m.Id != movie.Id && strings.EqualFold(strings.ToLower(m.Title), strings.ToLower(movie.Title)) {
						isExisting = true
						break
					}
				}
				if isExisting {
					WarningLogger.Printf("Error: Movie title %v already exists\n", movie.Title)
					http.Error(w, "Error: Movie title already exists. Please try another one.", http.StatusBadRequest)
					db.mu.Unlock()
					return
				}

				movieIndex = i
				movie.Id = item.Id
				movie.LastModified = time.Now().Format("2006-01-02 15:04:05")

				tempLog.Id = db.nextLogId
				tempLog.UserId = userId
				if item.Title != movie.Title {
					tempLog.Description = "Movie title was changed from " + item.Title + " to " + movie.Title
				} else if item.IsWatched != movie.IsWatched {
					if movie.IsWatched {
						tempLog.Description = "Movie " + movie.Title + " was marked as watched"
					} else {
						tempLog.Description = "Movie " + movie.Title + " was unmarked as watched"
					}
				} else {
					if movie.IsWatched {
						tempLog.Description = "Movie title was changed from " + item.Title + " to " + movie.Title + " and was marked as watched"
					} else {
						tempLog.Description = "Movie title was changed from " + item.Title + " to " + movie.Title + " and was unmarked as watched"
					}
				}
				tempLog.Date = time.Now().Format("2006-01-02 15:04:05")
				db.nextLogId++
				db.logs = append(db.logs, tempLog)

				break
			}
		}
		if movieIndex >= 0 {
			db.movieList[movieIndex] = movie

			InfoLogger.Printf("Movie ID %v has been edited\n", movie.Id)
		}
		db.mu.Unlock()

		// Returns error if movie ID provided does not exists
		if movieIndex < 0 {
			ErrorLogger.Printf("Error: Failed to edit movie. Movie ID %v does not exist\n", id)
			http.Error(w, "Error: Failed to edit movie. Movie does not exists.", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of movies of a user if not error
		if err := json.NewEncoder(w).Encode(filterMoviesByUser(db, userId)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		var userId int
		var tempLog Log
		var isExisting = false

		// Deletes a movie of a user and adds a log to the database
		db.mu.Lock()
		for j, movie := range db.movieList {
			if id == movie.Id {
				isExisting = true
				userId = movie.UserId

				tempLog.Id = db.nextLogId
				tempLog.UserId = userId
				tempLog.Description = "Movie " + movie.Title + " was deleted from the movie list"
				tempLog.Date = time.Now().Format("2006-01-02 15:04:05")
				db.nextLogId++
				db.logs = append(db.logs, tempLog)

				db.movieList = append(db.movieList[:j], db.movieList[j+1:]...)

				break
			}
		}
		db.mu.Unlock()

		// Returns error if movie ID provided does not exists
		if !isExisting {
			ErrorLogger.Printf("Error: Failed to delete movie. Movie ID %v does not exist\n", id)
			http.Error(w, "Error: Failed to delete movie. Movie does not exists.", http.StatusBadRequest)
			return
		}

		InfoLogger.Printf("Movie ID %v has been deleted\n", id)

		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of movies of a user if not error
		if err := json.NewEncoder(w).Encode(filterMoviesByUser(db, userId)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (db *Database) processLogs(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var tempLog Log

		// Returns error if request body does not fit with the log structure
		if err := json.NewDecoder(r.Body).Decode(&tempLog); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.mu.Lock()
		tempLog.Id = db.nextLogId
		tempLog.UserId = id
		tempLog.Date = time.Now().Format("2006-01-02 15:04:05")
		db.nextLogId++
		db.logs = append(db.logs, tempLog)

		if tempLog.Description == "User logged in" {
			InfoLogger.Printf("User ID %v logged in\n", id)
		}

		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of logs of a user if not error
		if err := json.NewEncoder(w).Encode(filterLogsByUser(db, id)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		// Returns updated list of logs of a user if not error
		if err := json.NewEncoder(w).Encode(filterLogsByUser(db, id)); err != nil {
			ErrorLogger.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// filterMoviesByUser returns list of movies of a user
func filterMoviesByUser(db *Database, id int) []Movie {
	filteredMovies := []Movie{}

	for _, movie := range db.movieList {
		if movie.UserId == id {
			filteredMovies = append(filteredMovies, movie)
		}
	}

	filteredMovies = sortMovies(filteredMovies)

	return filteredMovies
}

// sortMovies returns sorted list of movies of a user
func sortMovies(movieList []Movie) []Movie {
	// Sorts the movies by date added in ascending order
	sort.SliceStable(movieList, func(i, j int) bool {
		return movieList[i].DateAdded < movieList[j].DateAdded
	})

	// Sorts the movies if a user already watched it or not
	sort.SliceStable(movieList, func(i, j int) bool {
		return strconv.FormatBool(movieList[i].IsWatched) < strconv.FormatBool(movieList[j].IsWatched)
	})

	return movieList
}

// filterLogsByUser returns list of logs of a user
func filterLogsByUser(db *Database, id int) []Log {
	filteredLogs := []Log{}

	for _, log := range db.logs {
		if log.UserId == id {
			filteredLogs = append(filteredLogs, log)
		}
	}

	// Sorts the logs by date in descending order
	sort.SliceStable(filteredLogs, func(i, j int) bool {
		return filteredLogs[i].Date > filteredLogs[j].Date
	})

	return filteredLogs
}
