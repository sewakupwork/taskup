package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"taskup/models"
	"time"
)

func (rs taskRouter) Create(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var tsk models.Task
	if err := json.Unmarshal(body, &tsk); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("%v", err.Error()))
		return
	}

	// validate the basic fields
	if err := tsk.CreateRequestValidate(); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("%v", err.Error()))
		return
	}

	now := time.Now().UTC()

	tsk.CreatedAt = now
	tsk.UpdatedAt = now

	// Append to the Tasks table
	if err := rs.rt.DbConn.CreateTask(&tsk); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fmt.Sprintf("Task Created ID = %v", tsk.ID))
}
