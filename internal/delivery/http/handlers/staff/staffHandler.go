package staff

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

func StaffRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/", getAllStaff).Methods("GET")
}

func getAllStaff(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Staff : To be implemented"))
	if err != nil {
		return
	}
}
