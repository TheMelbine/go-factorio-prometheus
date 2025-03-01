package meters

const research_cmd = `/silent-command
local lines = {};
local force = game.forces["player"];
if (force and force.current_research) then
	table.insert(lines, string.format("%s,%s", force.research_progress, force.current_research.name))
end
rcon.print(table.concat(lines, "\n"))`

func ResearchMeters(manager *Manager) {
	manager.NewGaugeInt64(
		"current_research_progress",
		"The number of players that have joined the server",
		[]string{"technology"},
		CSVScraper[int64](
			"amount,technology",
			research_cmd,
		),
	)
}
