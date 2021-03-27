package users

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
}

var usersGroup []User

func (c Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(usersGroup)
}

func response(w http.ResponseWriter, data *User, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(data)
}

func (c Controller) PostUser(w http.ResponseWriter, r *http.Request) {
	var userAux User
	err := json.NewDecoder(r.Body).Decode(&userAux)
	if err != nil {
		response(w, &userAux, http.StatusInternalServerError)
		return
	}
	usersGroup = append(usersGroup, userAux)
	response(w, &userAux, http.StatusOK)
}
