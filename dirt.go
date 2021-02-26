package main

import (
	"github.com/StevenLeRoux/dirt/pkg/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Could not execute netrig")
	}
}
