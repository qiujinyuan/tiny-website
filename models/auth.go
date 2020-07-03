package models

import uuid "github.com/satori/go.uuid"

// Auth ...
type Auth struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

// CheckAuth ...
func CheckAuth(username, password string) bool {
	notFound := db.Select("id").Where(Auth{Username: username, Password: password}).RecordNotFound()
	if notFound {
		return false
	}
	return true
}
