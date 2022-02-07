package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	DateAdded string
}

type Movie struct {
	Id           int
	UserId       int
	Title        string
	IsWatched    bool
	DateAdded    string
	LastModified string
}

type Log struct {
	Id          int
	UserId      int
	Description string
	Date        string
}

type Database struct {
	nextUserId  int
	nextMovieId int
	nextLogId   int
	mu          sync.Mutex
	users       []User
	movieList   []Movie
	logs        []Log
}

func main() {
	db := &Database{users: []User{}, movieList: []Movie{}, logs: []Log{}}
	db.nextUserId = 1
	db.nextMovieId = 1
	db.nextLogId = 1
	http.ListenAndServe(":8080", db.handler())
}

func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		var id int

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
			log.Fatalln("Incorrect URL.")
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (db *Database) processUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user User
		var log Log

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.mu.Lock()

		user.Id = db.nextUserId
		user.DateAdded = time.Now().Format("2006-01-02 15:04:05")
		db.nextUserId++
		db.users = append(db.users, user)

		log.Id = db.nextLogId
		log.UserId = user.Id
		log.Description = "Account was created"
		log.Date = time.Now().Format("2006-01-02 15:04:05")
		db.nextLogId++
		db.logs = append(db.logs, log)

		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(db.users)
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(db.users)
	}
}

func (db *Database) processMovieList(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var movie Movie
		var log Log

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.mu.Lock()

		movie.Id = db.nextMovieId
		movie.UserId = id
		movie.DateAdded = time.Now().Format("2006-01-02 15:04:05")
		movie.LastModified = time.Now().Format("2006-01-02 15:04:05")
		db.nextMovieId++
		db.movieList = append(db.movieList, movie)

		log.Id = db.nextLogId
		log.UserId = id
		log.Description = "Movie " + movie.Title + " was added in the movie list"
		log.Date = time.Now().Format("2006-01-02 15:04:05")
		db.nextLogId++
		db.logs = append(db.logs, log)

		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(filterMoviesByUser(db, id))
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(filterMoviesByUser(db, id))
	case "PUT":
		var userId int
		var movie Movie
		var log Log

		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		movieIndex := -1

		db.mu.Lock()
		for i, item := range db.movieList {
			if id == item.Id {
				userId = item.UserId

				movieIndex = i
				movie.Id = item.Id
				movie.LastModified = time.Now().Format("2006-01-02 15:04:05")

				log.Id = db.nextLogId
				log.UserId = userId
				if item.Title != movie.Title {
					log.Description = "Movie title was changed from " + item.Title + " to " + movie.Title
				} else if item.IsWatched != movie.IsWatched {
					if movie.IsWatched {
						log.Description = "Movie " + movie.Title + " was marked as watched"
					} else {
						log.Description = "Movie " + movie.Title + " was unmarked as watched"
					}
				} else {
					if movie.IsWatched {
						log.Description = "Movie title was changed from " + item.Title + " to " + movie.Title + " and was marked as watched"
					} else {
						log.Description = "Movie title was changed from " + item.Title + " to " + movie.Title + " and was unmarked as watched"
					}
				}
				log.Date = time.Now().Format("2006-01-02 15:04:05")
				db.nextLogId++
				db.logs = append(db.logs, log)

				break
			}
		}
		if movieIndex >= 0 {
			db.movieList[movieIndex] = movie
		}
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(filterMoviesByUser(db, userId))
	case "DELETE":
		var userId int
		var log Log

		db.mu.Lock()
		for j, movie := range db.movieList {
			if id == movie.Id {
				userId = movie.UserId

				log.Id = db.nextLogId
				log.UserId = userId
				log.Description = "Movie " + movie.Title + " was deleted from the movie list"
				log.Date = time.Now().Format("2006-01-02 15:04:05")
				db.nextLogId++
				db.logs = append(db.logs, log)

				db.movieList = append(db.movieList[:j], db.movieList[j+1:]...)

				break
			}
		}
		db.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(filterMoviesByUser(db, userId))
	}
}

func (db *Database) processLogs(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")

		filteredLogs := []Log{}

		for _, log := range db.logs {
			if log.UserId == id {
				filteredLogs = append(filteredLogs, log)
			}
		}

		sort.SliceStable(filteredLogs, func(i, j int) bool {
			return filteredLogs[i].Date > filteredLogs[j].Date
		})

		json.NewEncoder(w).Encode(filteredLogs)
	}
}

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

func sortMovies(movieList []Movie) []Movie {
	sort.SliceStable(movieList, func(i, j int) bool {
		return movieList[i].DateAdded < movieList[j].DateAdded
	})

	sort.SliceStable(movieList, func(i, j int) bool {
		return strconv.FormatBool(movieList[i].IsWatched) < strconv.FormatBool(movieList[j].IsWatched)
	})

	return movieList
}
