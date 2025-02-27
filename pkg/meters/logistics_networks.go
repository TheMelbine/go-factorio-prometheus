package meters

// https://lua-api.factorio.com/latest/classes/LuaForce.html
const logistics_cmd = `/silent-command
local lines = {};
local networks = game.forces["player"].logistic_networks;
if force.current_research then
	table.insert(lines, string.format("%s,%s", force.research_progress, force.current_research.name))
end
rcon.print(table.concat(lines, "\n"))`

func LogisticsMeters(manager *Manager) {
	// manager.NewGaugeInt64(
	// 	"current_research_progress",
	// 	"The number of players that have joined the server",
	// 	[]string{"technology"},
	// 	CSVScraper[int64](
	// 		"amount,technology",
	// 		logistics_cmd,
	// 	),
	// )
}
