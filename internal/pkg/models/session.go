package models

import (
	uuid "github.com/satori/go.uuid"
)

const (
	SessionCookieName = "session_id"
	SessionContextKey = "session_key"
)

type Session struct {
	Value  string
	UserId uint64
}

func NewSession(userId uint64) *Session {
	newValue := uuid.NewV4()
	return &Session{
		Value:  newValue.String(),
		UserId: userId,
	}
}
