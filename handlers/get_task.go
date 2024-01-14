package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (rs taskRouter) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	tsk, err := rs.rt.DbConn.GetTask(id)
	if err != nil {
		fmt.Println("got error = ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tsk)
}
