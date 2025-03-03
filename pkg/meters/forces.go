package meters

import "github.com/daanv2/go-factorio-otel/pkg/meters/cost"

const force_cmd = `/silent-command
local lines = {};
local force = game.forces["player"];
if force then
	table.insert(lines, string.format("1,%s,%s,%s", force.name, force.research_progress, force.current_research.name, force.rockets_launched));
end
rcon.print(table.concat(lines, "\n"))`

func ForcesMeters(manager *Manager) {
	force_table := NewCSVTable(
		"force_table",
		force_cmd,
		[]string{"amount", "name", "research_progress", "current_research", "rockets_launched"},
	)
	manager.AddMeter(force_table)

	// Force
	manager.NewGaugeInt64(
		"forces",
		"List of all tables",
		[]string{"name"},
		force_table.SubTableInt64("amount", "name"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"forces_rockets_launched",
		"The number of rockets launched.",
		[]string{"name"},
		force_table.SubTableInt64("rockets_launched", "name"),
	).SetCost(cost.NONE)

	// Research
	manager.NewGaugeFloat64(
		"research_current_progress",
		"Progress of current research, as a number in range [0, 1].",
		[]string{"current_research", "name"},
		force_table.SubTableFloat64("research_progress", "current_research", "name"),
	).SetCost(cost.NONE)
}
