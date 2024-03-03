<p align="center">
  <img src="images/puregym.png" />
</p>

<h1 align="center">
    <a href="https://github.com/jackcoble/puregym-go">puregym-go</a>
</h1>
<p align="center">An API client for PureGym written in Go 💪</p>

# Getting Started

## Installation

```bash
$ go get -u github.com/jackcoble/puregym-go
```

## Usage Example

The code snippet below shows you how to use the PureGym API client and fetch the current number of people within your Home Gym.

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/jackcoble/puregym-go"
)

func main() {
	ctx := context.Background()

	// Create new instance of the PureGym API Client
	pureGym, err := puregym.NewClient(os.Getenv("PUREGYM_EMAIL"), os.Getenv("PUREGYM_PIN"))
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}

	// Authenticate against the API
	if err := pureGym.Authenticate(ctx); err != nil {
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

	log.Printf("Total number of people inside the Gym: %d.", gymAttendance.TotalPeopleInGym)
}
```

## License

This project is licensed under the terms of the [MIT License](https://opensource.org/license/mit).
