package database

import (
	"database/sql"
	"errors"

	"self-improvement/zoom-clone-backend/config"
	"self-improvement/zoom-clone-backend/models"
)

// SaveRoom inserts a new room into the database
func SaveRoom(room models.Room) error {
	query := `INSERT INTO rooms (host_id, name, password) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(query, room.HostID, room.Name, room.Password)
	if err != nil {
		return err
	}
	return nil
}

// GetRoomByID retrieves a room by its ID
func GetRoomByID(roomID string) (models.Room, error) {
	query := `SELECT id, host_id, name, password FROM rooms WHERE id = $1`
	row := config.DB.QueryRow(query, roomID)

	var room models.Room
	err := row.Scan(&room.ID, &room.HostID, &room.Name, &room.Password)
	if err == sql.ErrNoRows {
		return models.Room{}, errors.New("room not found")
	}
	if err != nil {
		return models.Room{}, err
	}

	return room, nil
}
