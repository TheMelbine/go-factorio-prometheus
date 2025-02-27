package meters

func surfaceProperty(property string) string {
	return `/silent-command
local lines = {};
for _, surface in pairs(game.surfaces) do
  table.insert(lines, string.format("%s,%s", surface.` + property + `, surface.name))
end
rcon.print(table.concat(lines, "\n"))`
}

func PlanetsMeters(manager *Manager) {
	manager.NewGaugeFloat64(
		"planet_pollution_total",
		"The total pollution on the planet",
		[]string{"planet"},
		CSVScraper[float64](
			"amount,planet",
			surfaceProperty("get_total_pollution()"),
		),
	)
	manager.NewGaugeFloat64(
		"planet_daytime",
		"Current time of day, as a number in range [0, 1).",
		[]string{"planet"},
		CSVScraper[float64](
			"amount,planet",
			surfaceProperty("daytime"),
		),
	)
	manager.NewGaugeFloat64(
		"planet_darkness",
		"Amount of darkness at the current time, as a number in range [0, 1].",
		[]string{"planet"},
		CSVScraper[float64](
			"amount,planet",
			surfaceProperty("darkness"),
		),
	)
	manager.NewGaugeFloat64(
		"planet_wind_speed",
		"Current wind speed in tiles per tick.",
		[]string{"planet"},
		CSVScraper[float64](
			"amount,planet",
			surfaceProperty("wind_speed"),
		),
	)
}
