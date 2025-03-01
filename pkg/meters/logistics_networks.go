package meters

// https://lua-api.factorio.com/latest/classes/LuaForce.html
const logistics_cmd = `/silent-command
local lines = {};
local force = game.forces["player"];
if (force and force.logistic_networks) then
	table.insert(lines, string.format("1,%s,%s,%s", force.name, force.research_progress, force.current_research.name, force.rockets_launched))
end
rcon.print(table.concat(lines, "\n"))`

func LogisticsMeters(manager *Manager) {
	logistics_table := NewCSVTable(
		"logistic_table",
		logistics_cmd,
		[]string{"amount", "name", "research_progress", "current_research", "rockets_launched"},
	)

	manager.AddMeter(logistics_table)
}
