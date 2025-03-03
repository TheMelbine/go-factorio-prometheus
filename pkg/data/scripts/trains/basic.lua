/silent-command
local lines = {};
local trains = game.train_manager.get_trains({})
for _, train in ipairs(trains) do
    table.insert(lines, string.format("%s,%s,%s,%s,%s,%s,%s,%s",
    train.speed,
    train.weight,
    train.get_item_count(nil),
    train.get_fluid_count(nil),
    train.kill_count,
    train.has_path,
    train.id,
    train.state
    ));
end
rcon.print(table.concat(lines, "\n"))