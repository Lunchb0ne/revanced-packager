package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/lunchb0ne/revanced-packager/internal/utils"
)

func main() {
	var log_level log.Level

	if os.Getenv("DEBUG") == "true" {
		log_level = log.DebugLevel
	} else {
		log_level = log.InfoLevel
	}

	logger := log.Logger{
		Handler: cli.New(os.Stdout),
		Level:   log_level,
	}

	preflight := logger.WithFields(log.Fields{
		"Type":  "Preflight",
		"Level": log_level,
	})

	preflight.Debug("Running Java Preflight")

	// check if we have the requirements
	err := utils.JavaPreflightCheck()

	if err != nil {
		// quit if java not found
		preflight.WithError(err).Error("Java Preflight Check Failed")
		return
	}

	preflight.Debug("Java Preflight Check Passed")

	// select the app you want
	// setup an array of strings to select the apps
	// apps := []string{"Youtube", "YouTube Music", "Twitter"}

}
