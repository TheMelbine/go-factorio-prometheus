package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
)

//go:embed scripts/players/basic.lua
var player_basic_cmd string

func PlayerMeters(manager *meters.Manager) {
	players_table := meters.NewCSVTable(
		"players_table",
		player_basic_cmd,
		[]string{"amount", "name", "planet", "connected", "afk_time", "online_time", "character_running_speed", "in_combat"},
	)
	manager.AddMeter(players_table)

	manager.NewGaugeInt64(
		"players_connected",
		"The state of players connected",
		[]string{"name", "planet", "connected", "in_combat"},
		players_table.SubTableInt64("amount", "name", "planet", "connected", "in_combat"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"players_afk_time",
		"The amount of time the players is afk",
		[]string{"name", "planet", "connected", "in_combat"},
		players_table.SubTableInt64("afk_time", "name", "planet", "connected", "in_combat"),
	).SetCost(cost.NONE)
	manager.NewGaugeInt64(
		"players_online_time",
		"The state of players connected",
		[]string{"name", "planet", "connected", "in_combat"},
		players_table.SubTableInt64("online_time", "name", "planet", "connected", "in_combat"),
	).SetCost(cost.NONE)
	manager.NewGaugeFloat64(
		"players_running_speed",
		"The current movement speed of this character, including effects from exoskeletons, tiles, stickers and shooting.",
		[]string{"name", "planet", "connected", "in_combat"},
		players_table.SubTableFloat64("character_running_speed", "name", "planet", "connected", "in_combat"),
	).SetCost(cost.NONE)
}
