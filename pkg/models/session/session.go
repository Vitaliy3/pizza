package session

import (
	"agile/pkg/models/auth"

	"github.com/google/uuid"
)

var Sessions = make(map[string]auth.User)

func Add(user auth.User) auth.User {
	newtoken, _ := uuid.NewUUID()
	user.AccessToken = newtoken.String()
	Sessions[newtoken.String()] = user

	return user
}

func Remove(token string) {
	delete(Sessions, token)
}
