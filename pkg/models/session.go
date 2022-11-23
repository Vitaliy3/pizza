package models

import (
	"github.com/google/uuid"
)

var Sessions = make(map[string]User)

func Add(user User) User {
	newtoken, _ := uuid.NewUUID()
	user.AccessToken = newtoken.String()
	Sessions[newtoken.String()] = user

	return user
}

func Remove(token string) {
	delete(Sessions, token)
}
