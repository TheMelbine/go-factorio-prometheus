package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
)

//go:embed scripts/platforms/basic.lua
var platforms_basic_cmd string

func PlatformMeters(manager *meters.Manager) {
	platforms_table := meters.NewCSVTable(
		"platforms_table",
		platforms_basic_cmd,
		[]string{
			"speed",
			"weight",
			"name",
			"state",
			"paused",
			"force",
		},
	)
	manager.AddMeter(platforms_table)

	manager.NewGaugeFloat64(
		"platforms_speed",
		"Current speed",
		[]string{"name", "state", "paused", "force"},
		platforms_table.SubTableFloat64("speed", "name", "state", "paused", "force"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"platforms_weight",
		"The weight of platforms",
		[]string{"name", "state", "paused", "force"},
		platforms_table.SubTableFloat64("weight", "name", "state", "paused", "force"),
	).SetCost(cost.NONE)
}
