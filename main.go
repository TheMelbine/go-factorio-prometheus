package main

import (
	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-prometheus/cmd"
	"github.com/daanv2/go-factorio-prometheus/internal/setup"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	setup.Logger()
	_, err := maxprocs.Set(maxprocs.Logger(log.Infof))
	if err != nil {
		log.Fatal("error setting maxprocs", "error", err)
	}
}

func main() {
	cmd.Execute()
}
