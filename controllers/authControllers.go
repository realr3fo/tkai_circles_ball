package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/realr3fo/tkai_circles_ball/models"
	u "github.com/realr3fo/tkai_circles_ball/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, err := account.Create() //Create account
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp, err := models.Login(account.Username, account.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, resp)
}
