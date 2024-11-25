package models

type Room struct {
	ID       string `json:"id"`
	HostID   string `json:"host_id"`
	Name     string `json:"name"`
	Password string `json:"password,omitempty"` // Optional for private rooms
}
