package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"taskup/models"
	"time"
)

func (rs taskRouter) Update(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	id := chi.URLParam(r, "id")

	var chTsk models.Task
	json.Unmarshal(body, &chTsk)

	dbTsk, err := rs.rt.DbConn.GetTask(id)
	if err != nil {
		fmt.Println("got error = ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// validate the basic fields
	if err := chTsk.UpdateRequestValidate(); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("%v", err.Error()))
		return
	}

	if chTsk.Description != "" {
		dbTsk.Description = chTsk.Description
	}

	if !chTsk.DueDate.IsZero() {
		dbTsk.DueDate = chTsk.DueDate
	}

	if chTsk.Priority != "" {
		dbTsk.Priority = chTsk.Priority
	}

	dbTsk.UpdatedAt = time.Now().UTC()

	if err := rs.rt.DbConn.UpdateTask(dbTsk); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Updated")

}
