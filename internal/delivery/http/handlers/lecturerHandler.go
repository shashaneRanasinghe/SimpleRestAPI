package handlers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func LecturerRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/", getAllLecturers).Methods("GET")
}

func getAllLecturers(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Lecturer : To be implemented"))
	if err != nil {
		return
	}
}
