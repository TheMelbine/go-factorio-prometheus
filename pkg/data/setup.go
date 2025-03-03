package data

import "github.com/daanv2/go-factorio-otel/pkg/meters"

func Setup(manager *meters.Manager) {
	PlayerMeters(manager)
	PlanetsMeters(manager)
	ForcesMeters(manager)
	LogisticsMeters(manager)
}
