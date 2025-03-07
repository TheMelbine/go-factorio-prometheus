package data

import (
	_ "embed"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
)

var (
	//go:embed scripts/production/items-input.lua
	items_input_cmd string
	//go:embed scripts/production/items-output.lua
	items_output_cmd string
	//go:embed scripts/production/items-storage.lua
	items_storage_cmd string

	//go:embed scripts/production/fluids-input.lua
	fluids_input_cmd string
	//go:embed scripts/production/fluids-output.lua
	fluids_output_cmd string
	//go:embed scripts/production/fluids-storage.lua
	fluids_storage_cmd string

	//go:embed scripts/production/kills-input.lua
	kills_input_cmd string
	//go:embed scripts/production/kills-output.lua
	kills_output_cmd string
	//go:embed scripts/production/kills-storage.lua
	kills_storage_cmd string

	//go:embed scripts/production/pollution-input.lua
	pollution_input_cmd string
	//go:embed scripts/production/pollution-output.lua
	pollution_output_cmd string
	//go:embed scripts/production/pollution-storage.lua
	pollution_storage_cmd string

	//go:embed scripts/production/build-input.lua
	entity_build_input_cmd string
	//go:embed scripts/production/build-output.lua
	entity_build_output_cmd string
	//go:embed scripts/production/build-storage.lua
	entity_build_storage_cmd string

	//go:embed scripts/production/electric-input.lua
	electric_input_cmd string
	//go:embed scripts/production/electric-output.lua
	electric_output_cmd string
	//go:embed scripts/production/electric-storage.lua
	electric_storage_cmd string
)

func ProductionMeters(manager *meters.Manager) {
	// Items
	{
		manager.NewGaugeInt64(
			"production_items_input",
			"Amount of items being inputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				items_input_cmd,
			),
		)
		manager.NewGaugeInt64(
			"production_items_output",
			"Amount of items being outputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				items_output_cmd,
			),
		)
		manager.NewGaugeInt64(
			"production_items_storage",
			"Amount of items in storage",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				items_storage_cmd,
			),
		)
	}
	// Fluids
	{
		manager.NewGaugeFloat64(
			"production_fluids_input",
			"Amount of fluids being inputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				fluids_input_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_fluids_output",
			"Amount of fluids being outputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				fluids_output_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_fluids_storage",
			"Amount of fluids in storage",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				fluids_storage_cmd,
			),
		)
	}
	// Kills
	{
		manager.NewGaugeInt64(
			"production_kills_input",
			"Amount of kills being inputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				kills_input_cmd,
			),
		)
		manager.NewGaugeInt64(
			"production_kills_output",
			"Amount of kills being outputted",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				kills_output_cmd,
			),
		)
		manager.NewGaugeInt64(
			"production_kills_storage",
			"Amount of kills in storage",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[int64](
				"amount,name,planet,force",
				kills_storage_cmd,
			),
		)
	}
	// Pollution
	{
		manager.NewGaugeFloat64(
			"production_pollution_input",
			"Amount of pollution being inputted",
			[]string{"name", "planet"},
			meters.CSVScraper[float64](
				"amount,name,planet",
				pollution_input_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_pollution_output",
			"Amount of pollution being outputted",
			[]string{"name", "planet"},
			meters.CSVScraper[float64](
				"amount,name,planet",
				pollution_output_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_pollution_storage",
			"Amount of pollution in storage",
			[]string{"name", "planet"},
			meters.CSVScraper[float64](
				"amount,name,planet",
				pollution_storage_cmd,
			),
		)
	}
	// Entity build
	{
		manager.NewGaugeFloat64(
			"production_entity_build_input",
			"The entity build statistics for this force (built and mined) for the given surface.",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				entity_build_input_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_entity_build_output",
			"The entity build statistics for this force (built and mined) for the given surface.",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				entity_build_output_cmd,
			),
		)
		manager.NewGaugeFloat64(
			"production_entity_build_storage",
			"The entity build statistics for this force (built and mined) for the given surface.",
			[]string{"name", "planet", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,force",
				entity_build_storage_cmd,
			),
		)
	}
	// Electric build
	{
		manager.NewCounterFloat64(
			"production_electric_intput",
			"The electric network statistics for this electric pole. [joules]",
			[]string{"name", "planet", "network", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,network,force",
				electric_input_cmd,
			),
		)
		manager.NewCounterFloat64(
			"production_electric_output",
			"The electric network statistics for this electric pole. [joules]",
			[]string{"name", "planet", "network", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,network,force",
				electric_output_cmd,
			),
		)
		manager.NewCounterFloat64(
			"production_electric_storage",
			"The electric network statistics for this electric pole. [joules]",
			[]string{"name", "planet", "network", "force"},
			meters.CSVScraper[float64](
				"amount,name,planet,network,force",
				electric_storage_cmd,
			),
		)
	}
}
