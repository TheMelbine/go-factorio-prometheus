package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
)

//go:embed scripts/trains/basic.lua
var trains_basic_cmd string

func TrainMeters(manager *meters.Manager) {
	trains_table := meters.NewCSVTable(
		"trains_table",
		trains_basic_cmd,
		[]string{
			"speed",
			"weight",
			"item_count",
			"fluid_count",
			"kill_count",
			"has_path",
			"id",
			"state",
		},
	)
	manager.AddMeter(trains_table)

	manager.NewGaugeFloat64(
		"trains_speed",
		"Current speed",
		[]string{"has_path", "id", "state"},
		trains_table.SubTableFloat64("speed", "has_path", "id", "state"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"trains_weight",
		"The weight of trains",
		[]string{"has_path", "id", "state"},
		trains_table.SubTableFloat64("weight", "has_path", "id", "state"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"trains_item_count",
		"The amount of a particular item stored in the train.",
		[]string{"has_path", "id", "state"},
		trains_table.SubTableInt64("item_count", "has_path", "id", "state"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"trains_fluid_count",
		"The amount of a particular fluid stored in the train.",
		[]string{"has_path", "id", "state"},
		trains_table.SubTableInt64("fluid_count", "has_path", "id", "state"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"trains_kill_count",
		"The total number of kills by this train.",
		[]string{"has_path", "id", "state"},
		trains_table.SubTableInt64("kill_count", "has_path", "id", "state"),
	).SetCost(cost.NONE)
}
