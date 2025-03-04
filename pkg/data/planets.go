package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
)

//go:embed scripts/planets/basic.lua
var planets_table_cmd string

func PlanetsMeters(manager *meters.Manager) {
	planet_table := meters.NewCSVTable(
		"planets_table",
		planets_table_cmd,
		[]string{"planet", "total_pollution", "daytime", "darkness", "wind_speed"},
	)
	manager.AddMeter(planet_table)

	manager.NewGaugeFloat64(
		"planet_pollution_total",
		"The total pollution on the planet",
		[]string{"planet"},
		planet_table.SubTableFloat64("total_pollution", "planet"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"planet_daytime",
		"Current time of day, as a number in range [0, 1).",
		[]string{"planet"},
		planet_table.SubTableFloat64("daytime", "planet"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"planet_darkness",
		"Amount of darkness at the current time, as a number in range [0, 1].",
		[]string{"planet"},
		planet_table.SubTableFloat64("darkness", "planet"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"planet_wind_speed",
		"Amount of darkness at the current time, as a number in range [0, 1].",
		[]string{"planet"},
		planet_table.SubTableFloat64("wind_speed", "planet"),
	).SetCost(cost.NONE)
}
