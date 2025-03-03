package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-otel/pkg/meters"
	"github.com/daanv2/go-factorio-otel/pkg/meters/cost"
)

// https://lua-api.factorio.com/latest/classes/LuaForce.html

//go:embed scripts/logistics/robots.lua
var logistics_robots_cmd string

func LogisticsMeters(manager *meters.Manager) {
	logistics_robots_table := meters.NewCSVTable(
		"logistics_robots_table",
		logistics_robots_cmd,
		[]string{"available_logistic_robots", "all_logistic_robots", "available_construction_robots", "all_construction_robots", "network_id"},
	)
	manager.AddMeter(logistics_robots_table)

	manager.NewGaugeInt64(
		"logistics_available_logistic_robots",
		"The amount of available logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("available_logistic_robots", "network_id"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"logistics_all_logistic_robots",
		"The amount of total amount of logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("all_logistic_robots", "network_id"),
	).SetCost(cost.NONE)
	
	manager.NewGaugeInt64(
		"logistics_available_construction_robots",
		"The amount of available logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("available_construction_robots", "network_id"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"logistics_all_construction_robots",
		"The amount of total amount of logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("all_construction_robots", "network_id"),
	).SetCost(cost.NONE)
}
