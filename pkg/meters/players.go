package meters

const player_joined_cmd = `/silent-command
local lines = {};
for _, player in pairs(game.connected_players) do
    table.insert(lines, string.format("1,%s,%s,%s",
        player.name,
        player.surface.name,
        tostring(player.connected)
    ))
end
rcon.print(table.concat(lines, "\n"))`

func PlayerMeters(manager *Manager) {
	manager.NewGaugeInt64(
		"players_joined",
		"The number of players that have joined the server",
		[]string{"name", "planet", "connected"},
		CSVScraper[int64](
			"amount,name,planet,connected",
			force_joined_cmd,
		),
	)
}
