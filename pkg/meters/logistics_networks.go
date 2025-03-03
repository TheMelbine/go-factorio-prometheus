package meters

import "github.com/daanv2/go-factorio-otel/pkg/meters/cost"

// https://lua-api.factorio.com/latest/classes/LuaForce.html
const logistics_robots_cmd = `/silent-command
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
rcon.print(table.concat(lines, "\n"))`

func LogisticsMeters(manager *Manager) {
	logistics_robots_table := NewCSVTable(
		"logistics_robots_table",
		logistics_robots_cmd,
		[]string{"available_logistic_robots", "all_logistic_robots", "available_construction_robots", "all_construction_robots", "network_id"},
	)
	manager.AddMeter(logistics_robots_table)

	manager.NewGaugeInt64(
		"logistics_available_logistic_robots",
		"The amount of available logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("available_logistic_robots", "network_id"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"logistics_all_logistic_robots",
		"The amount of total amount of logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("all_logistic_robots", "network_id"),
	).SetCost(cost.NONE)
	
	manager.NewGaugeInt64(
		"logistics_available_construction_robots",
		"The amount of available logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("available_construction_robots", "network_id"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"logistics_all_construction_robots",
		"The amount of total amount of logistic robots",
		[]string{"network_id"},
		logistics_robots_table.SubTableInt64("all_construction_robots", "network_id"),
	).SetCost(cost.NONE)
}
