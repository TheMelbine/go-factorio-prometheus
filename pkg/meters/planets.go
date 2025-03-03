package meters

import "github.com/daanv2/go-factorio-otel/pkg/meters/cost"

const planets_table_cmd = `/silent-command
local lines = {};
for _, surface in pairs(game.surfaces) do
  table.insert(lines, string.format("%s,%s,%s,%s,%s", surface.name, surface.get_total_pollution(), surface.daytime, surface.darkness, surface.wind_speed));
end
rcon.print(table.concat(lines, "\n"))`

func PlanetsMeters(manager *Manager) {
	planet_table := NewCSVTable(
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
