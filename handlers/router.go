package handlers

import (
	"github.com/go-chi/chi/v5"
	"taskup/runtime"
)

type Router interface {
	// Routes creates a REST router for task
	Routes() chi.Router
}

func NewTaskRouter(r *runtime.Runtime) Router {
	return taskRouter{rt: r}
}

type taskRouter struct {
	rt *runtime.Runtime
}

func (rs taskRouter) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.List)    // GET /tasks - read a list of tasks
	r.Post("/", rs.Create) // POST /tasks - create a new tasks and persist it
	r.Delete("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)       // GET /tasks/{id} - read a single task by :id
		r.Patch("/", rs.Update)  // PUT /tasks/{id} - update a single task by :id
		r.Delete("/", rs.Delete) // DELETE /tasks/{id} - delete a single task by :id
	})

	return r
}
