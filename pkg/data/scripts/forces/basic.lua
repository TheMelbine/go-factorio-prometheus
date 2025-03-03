/silent-command
local lines = {};
local force = game.forces["player"];
if force then
	table.insert(lines, string.format("1,%s,%s,%s", force.name, force.research_progress, force.current_research.name, force.rockets_launched));
end
rcon.print(table.concat(lines, "\n"))