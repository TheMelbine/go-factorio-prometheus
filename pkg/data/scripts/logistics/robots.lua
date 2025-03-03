/silent-command
local lines = {};
local force = game.forces["player"];
if force then
	for _, networks in pairs(force.logistic_networks) do
		for _, network in ipairs(networks) do
			table.insert(lines, string.format("%s,%s,%s,%s,%s",
				network.available_logistic_robots,
				network.all_logistic_robots,
				network.available_construction_robots,
				network.all_construction_robots,
				network.network_id
			));
		end
	end
end
rcon.print(table.concat(lines, "\n"))