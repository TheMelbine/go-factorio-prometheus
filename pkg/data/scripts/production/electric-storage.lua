/silent-command
local lines = {}
local force = game.forces["player"]
if force then
  for _, surface in pairs(game.surfaces) do
    local processed_networks = {}
    local poles = surface.find_entities_filtered{force=force, type="electric-pole", limit=5}
    for _, pole in pairs(poles) do
      local net_id = pole.electric_network_id
      if net_id and not processed_networks[net_id] and pole.electric_network_statistics then
        processed_networks[net_id] = true
        local stats = pole.electric_network_statistics
        for k, v in pairs(stats.storage_counts) do
          table.insert(lines, string.format("%s,%s,%s,%s,%s", v, k, surface.name, net_id, force.name))
        end
      end
    end
  end
end
rcon.print(table.concat(lines, "\n"))
