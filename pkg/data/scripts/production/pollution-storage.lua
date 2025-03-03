/silent-command
local lines = {};
for _, surface in pairs(game.surfaces) do
    local stats = game.get_pollution_statistics(surface.index);
    for k, v in pairs(stats.storage_counts) do
        table.insert(lines, string.format("%s,%s,%s", v, k, surface.name))
    end
end
rcon.print(table.concat(lines, "\n"))
