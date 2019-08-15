package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../models"
	u "../utils"
)

//CreateMeetup is an action the creates a meetup and then sends the data tot eh model
var CreateMeetup = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	meetup := &models.Meetup{}

	err := json.NewDecoder(r.Body).Decode(meetup)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := meetup.Create()
	u.Respond(w, res)
}

//UpdateMeetup is an action the creates a meetup and then sends the data tot eh model
var UpdateMeetup = func(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// //id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	u.Respond(w, u.Message(false, "there was an error in your request"))
	// 	return
	// }

	meetup := &models.Meetup{}

	err := json.NewDecoder(r.Body).Decode(meetup)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := meetup.UpdateMeetup()
	u.Respond(w, res)
}

//DeleteMeetup sends id to model
var DeleteMeetup = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "there was an error in your request"))
		return
	}

	res := models.DeleteMeetup(uint(id))
	res["data"] = id
	u.Respond(w, res)
}

//ToggleLike sends id to model
var ToggleLike = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "there was an error in your request"))
		return
	}

	res := models.ToggleLike(uint(id))
	res["data"] = id
	u.Respond(w, res)
}

//GetMeetup gets 1 meetup based on id
var GetMeetup = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "there was an error in your request"))
		return
	}

	data := models.GetMeetup(uint(id))

	res := u.Message(true, "success")

	res["data"] = data

	u.Respond(w, res)

}

//GetMeetups gets all the meetups
var GetMeetups = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetMeetups()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
