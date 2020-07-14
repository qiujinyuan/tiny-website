package models

import uuid "github.com/satori/go.uuid"

// Auth ...
type Auth struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

// CheckAuth ...
func CheckAuth(username, password string) (exist bool, auth Auth) {
	notFound := db.Where(&Auth{Username: username, Password: password}).First(&auth).RecordNotFound()
	if notFound {
		exist = false
		return
	}
	exist = true
	return
}
