package services

import (
	"errors"
	database "self-improvement/zoom-clone-backend/DAO"
	"self-improvement/zoom-clone-backend/models"
)

// CreateRoom creates a new video conference room
func CreateRoom(data map[string]string) (models.Room, error) {
	if data["host_id"] == "" || data["name"] == "" {
		return models.Room{}, errors.New("host_id and name are required")
	}

	room := models.Room{
		HostID:   data["host_id"],
		Name:     data["name"],
		Password: data["password"], // Optional password
	}

	if err := database.SaveRoom(room); err != nil {
		return models.Room{}, err
	}

	return room, nil
}

// GetRoom retrieves a room by ID
func GetRoom(roomID string) (models.Room, error) {
	return database.GetRoomByID(roomID)
}
