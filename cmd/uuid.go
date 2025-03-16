package main

import "github.com/google/uuid"

func New() string {
	uuid := uuid.New()
	return uuid.String()
}
