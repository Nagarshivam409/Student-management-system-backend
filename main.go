package main

import (
	"backend/pkg/app"
	"backend/pkg/repository"
)

func main() {

	// Intial Migration
	repository.Intialmigration()
	// Intialising Router
	app.IntialiseRouter()
}
