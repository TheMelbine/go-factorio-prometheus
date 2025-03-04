/silent-command
local lines = {};
local force = game.forces["player"];
if force then
    for _, platform in pairs(force.platforms) do
        table.insert(lines, string.format("%s,%s,%s,%s,%s",
            platform.speed,
            platform.weight,
            platform.name,
            platform.state,
            platform.paused,
            force.name
        ));
    end
end
rcon.print(table.concat(lines, "\n"))