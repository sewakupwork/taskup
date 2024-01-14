package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (rs taskRouter) List(w http.ResponseWriter, r *http.Request) {
	taksList, err := rs.rt.DbConn.ListTask()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taksList)
}
