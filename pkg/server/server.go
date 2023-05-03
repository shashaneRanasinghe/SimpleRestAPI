package server

import (
	"context"
	"github.com/shashaneRanasinghe/simpleAPI/internal/delivery/http/handlers/lecturer"
	"github.com/shashaneRanasinghe/simpleAPI/internal/delivery/http/handlers/staff"
	"github.com/shashaneRanasinghe/simpleAPI/internal/delivery/http/handlers/student"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/tryfix/log"
)

// The Serve function creates the server
func Serve() chan string {
	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	db := database.NewDatabase()
	db.InitDatabase()
	conn := db.GetConnection()

	st := student.NewStudentHandler(conn)
	studentRouter := router.PathPrefix("/student").Subrouter()
	st.StudentRoutes(studentRouter)

	lecturerRouter := router.PathPrefix("/lecturer").Subrouter()
	lecturer.LecturerRoutes(lecturerRouter, conn)

	staffRouter := router.PathPrefix("/staff").Subrouter()
	staff.StaffRoutes(staffRouter, conn)

	closeChannel := make(chan string)

	//This goroutine will make sure that the service is stopped gracefully
	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)
		signal.Notify(sig, syscall.SIGTERM)
		signal.Notify(sig, syscall.SIGQUIT)
		<-sig

		log.Info("service interruption received")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Error("Server shutdown error : %v", err)
		}

		log.Info("HTTP server stopped")
		close(closeChannel)
	}()

	log.Info("server is starting on port " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}

	return closeChannel
}
