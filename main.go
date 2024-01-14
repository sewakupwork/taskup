package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"taskup/db"
	"taskup/handlers"
	"taskup/runtime"
)

func main() {
	// wait for db to be ready
	// time.Sleep(10 * time.Second)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	dbcli, err := db.NewDBClient()
	if err != nil {
		fmt.Println("failed to create db connection error = ", err)
		panic(err)
	}

	rt := &runtime.Runtime{
		DbConn: dbcli,
	}

	taskRoutes := handlers.NewTaskRouter(rt)
	r.Mount("/tasks", taskRoutes.Routes())
	fmt.Println("app is serving on port 8080")

	http.ListenAndServe(":8080", r)
}
