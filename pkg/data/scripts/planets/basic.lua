/silent-command
local lines = {};
for _, surface in pairs(game.surfaces) do
  table.insert(lines, string.format("%s,%s,%s,%s,%s", surface.name, surface.get_total_pollution(), surface.daytime, surface.darkness, surface.wind_speed));
end
rcon.print(table.concat(lines, "\n"))