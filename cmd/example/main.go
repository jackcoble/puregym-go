package main

import (
	"log"
	"os"

	"github.com/jackcoble/puregym-go"
)

func main() {
	// Create new instance of the PureGym API Client
	pureGym := puregym.NewClient(os.Getenv("PUREGYM_EMAIL"), os.Getenv("PUREGYM_PIN"))

	// Authenticate against the API
	if err := pureGym.Authenticate(); err != nil {
		log.Fatalf("unable to authenticate: %s", err.Error())
	}

	// Set the Home Gym for the client to use
	if err := pureGym.SetHomeGym(); err != nil {
		log.Printf("unable to set home gym for client: %s", err.Error())
	}

	// Get Live Capacity at Gym
	gymAttendance, err := pureGym.GetGymAttendance()
	if err != nil {
		log.Printf("unable to fetch live capacity: %s", err.Error())
	}

	log.Printf("Total number of people inside the Gym: %d", gymAttendance.TotalPeopleInGym)
}
