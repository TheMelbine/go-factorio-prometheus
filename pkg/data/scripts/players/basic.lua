/silent-command
local lines = {};
for _, player in pairs(game.players) do
    table.insert(lines, string.format("1,%s,%s,%s,%s,%s,%s",
        player.name,
        player.surface.name,
        player.connected,
        player.afk_time,
        player.online_time,
        player.character_running_speed,
        player.in_combat
    ));
end
rcon.print(table.concat(lines, "\n"))