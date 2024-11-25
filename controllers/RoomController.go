package controllers

import (
	"encoding/json"
	"net/http"
	"self-improvement/zoom-clone-backend/services"
	"self-improvement/zoom-clone-backend/utils"

	"github.com/gorilla/mux"
)

// CreateRoom handles room creation requests
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	room, err := services.CreateRoom(reqBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, room)
}

// GetRoom handles requests to retrieve room details
func GetRoom(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["id"]
	room, err := services.GetRoom(roomID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Room not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, room)
}
