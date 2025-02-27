package meters

const force_joined_cmd = `/silent-command
local lines = {};
for _, force in pairs(game.forces) do
    table.insert(lines, string.format("1,%s",
        force.name
    ))
end
rcon.print(table.concat(lines, "\n"))`

func ForcesMeters(manager *Manager) {
	manager.NewGaugeInt64(
		"forces",
		"The number of players that have joined the server",
		[]string{"force"},
		CSVScraper[int64](
			"amount,force",
			force_joined_cmd,
		),
	)
}
