package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aaronshurley/trackertix/tracker"
)

func main() {
	fmt.Println("Starting program...")

	token := os.Getenv("PIVOTAL_TRACKER_API_TOKEN")
	if token == "" {
		log.Fatal("PIVOTAL_TRACKER_API_TOKEN is not provided")
	}

	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("PROJECT_ID is not provided")
	}

	pid, err := strconv.Atoi(projectID)
	if err != nil {
		log.Fatal("PROJECT_ID is not a valid integer")
	}

	opts := tracker.ConfigOpts{
		Token:     token,
		ProjectID: pid,
	}
	tracker := tracker.NewTracker(opts)
	if err = tracker.SetupProject(); err != nil {
		log.Fatalf("SetupProject failed: %s", err)
	}
}
