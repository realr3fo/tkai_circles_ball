package controllers

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	"github.com/realr3fo/tkai_circles/models"
	u "github.com/realr3fo/tkai_circles/utils"
)

var GetCircleArea = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	if id == 0 {
		//The passed path parameter is not an integer
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	circle := &models.Circle{}
	err := json.NewDecoder(r.Body).Decode(circle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp, err := circle.GetCircleArea(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, resp)
}

var GetTubeVolume = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	if id == 0 {
		//The passed path parameter is not an integer
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	tube := &models.Tube{}
	err := json.NewDecoder(r.Body).Decode(tube)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp, err := tube.GetTubeVolume(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, resp)
}

var GetBallVolume = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	if id == 0 {
		//The passed path parameter is not an integer
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	ball := &models.Ball{}
	err := json.NewDecoder(r.Body).Decode(ball)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp, err := ball.GetBallVolume(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, err.Error()))
		return
	}
	u.Respond(w, resp)
}
