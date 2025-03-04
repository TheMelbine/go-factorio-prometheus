# Metrics

This document lists all metrics exposed by the application.

| Name | Help | Type |
|------|------|------|
| factorio_game_forces | List of all tables | GAUGE |
| factorio_game_forces_rockets_launched | The number of rockets launched. | GAUGE |
| factorio_game_logistics_all_construction_robots | The amount of total amount of logistic robots | GAUGE |
| factorio_game_logistics_all_logistic_robots | The amount of total amount of logistic robots | GAUGE |
| factorio_game_logistics_available_construction_robots | The amount of available logistic robots | GAUGE |
| factorio_game_logistics_available_logistic_robots | The amount of available logistic robots | GAUGE |
| factorio_game_planet_darkness | Amount of darkness at the current time, as a number in range [0, 1]. | GAUGE |
| factorio_game_planet_daytime | Current time of day, as a number in range [0, 1). | GAUGE |
| factorio_game_planet_pollution_total | The total pollution on the planet | GAUGE |
| factorio_game_planet_wind_speed | Amount of darkness at the current time, as a number in range [0, 1]. | GAUGE |
| factorio_game_players_afk_time | The amount of time the players is afk | GAUGE |
| factorio_game_players_connected | The state of players connected | GAUGE |
| factorio_game_players_online_time | The state of players connected | GAUGE |
| factorio_game_players_running_speed | The current movement speed of this character, including effects from exoskeletons, tiles, stickers and shooting. | GAUGE |
| factorio_game_production_fluids_input | Amount of fluids being inputted | GAUGE |
| factorio_game_production_fluids_output | Amount of fluids being outputted | GAUGE |
| factorio_game_production_fluids_storage | Amount of fluids in storage | GAUGE |
| factorio_game_production_items_input | Amount of items being inputted | GAUGE |
| factorio_game_production_items_output | Amount of items being outputted | GAUGE |
| factorio_game_production_items_storage | Amount of items in storage | GAUGE |
| factorio_game_production_kills_input | Amount of kills being inputted | GAUGE |
| factorio_game_production_kills_output | Amount of kills being outputted | GAUGE |
| factorio_game_production_kills_storage | Amount of kills in storage | GAUGE |
| factorio_game_production_pollution_input | Amount of pollution being inputted | GAUGE |
| factorio_game_production_pollution_output | Amount of pollution being outputted | GAUGE |
| factorio_game_production_pollution_storage | Amount of pollution in storage | GAUGE |
| factorio_game_research_current_progress | Progress of current research, as a number in range [0, 1]. | GAUGE |
| factorio_game_trains_fluid_count | The amount of a particular fluid stored in the train. | GAUGE |
| factorio_game_trains_item_count | The amount of a particular item stored in the train. | GAUGE |
| factorio_game_trains_kill_count | The total number of kills by this train. | GAUGE |
| factorio_game_trains_speed | Current speed | GAUGE |
| factorio_game_trains_weight | The weight of trains | GAUGE |
| go_gc_duration_seconds | A summary of the pause duration of garbage collection cycles. | SUMMARY |
| go_goroutines | Number of goroutines that currently exist. | GAUGE |
| go_info | Information about the Go environment. | GAUGE |
| go_memstats_alloc_bytes | Number of bytes allocated and still in use. | GAUGE |
| go_memstats_alloc_bytes_total | Total number of bytes allocated, even if freed. | COUNTER |
| go_memstats_buck_hash_sys_bytes | Number of bytes used by the profiling bucket hash table. | GAUGE |
| go_memstats_frees_total | Total number of frees. | COUNTER |
| go_memstats_gc_cpu_fraction | The fraction of this program's available CPU time used by the GC since the program started. | GAUGE |
| go_memstats_gc_sys_bytes | Number of bytes used for garbage collection system metadata. | GAUGE |
| go_memstats_heap_alloc_bytes | Number of heap bytes allocated and still in use. | GAUGE |
| go_memstats_heap_idle_bytes | Number of heap bytes waiting to be used. | GAUGE |
| go_memstats_heap_inuse_bytes | Number of heap bytes that are in use. | GAUGE |
| go_memstats_heap_objects | Number of allocated objects. | GAUGE |
| go_memstats_heap_released_bytes | Number of heap bytes released to OS. | GAUGE |
| go_memstats_heap_sys_bytes | Number of heap bytes obtained from system. | GAUGE |
| go_memstats_last_gc_time_seconds | Number of seconds since 1970 of last garbage collection. | GAUGE |
| go_memstats_lookups_total | Total number of pointer lookups. | COUNTER |
| go_memstats_mallocs_total | Total number of mallocs. | COUNTER |
| go_memstats_mcache_inuse_bytes | Number of bytes in use by mcache structures. | GAUGE |
| go_memstats_mcache_sys_bytes | Number of bytes used for mcache structures obtained from system. | GAUGE |
| go_memstats_mspan_inuse_bytes | Number of bytes in use by mspan structures. | GAUGE |
| go_memstats_mspan_sys_bytes | Number of bytes used for mspan structures obtained from system. | GAUGE |
| go_memstats_next_gc_bytes | Number of heap bytes when next garbage collection will take place. | GAUGE |
| go_memstats_other_sys_bytes | Number of bytes used for other system allocations. | GAUGE |
| go_memstats_stack_inuse_bytes | Number of bytes in use by the stack allocator. | GAUGE |
| go_memstats_stack_sys_bytes | Number of bytes obtained from system for stack allocator. | GAUGE |
| go_memstats_sys_bytes | Number of bytes obtained from system. | GAUGE |
| go_threads | Number of OS threads created. | GAUGE |
| process_cpu_seconds_total | Total user and system CPU time spent in seconds. | COUNTER |
| process_max_fds | Maximum number of open file descriptors. | GAUGE |
| process_open_fds | Number of open file descriptors. | GAUGE |
| process_resident_memory_bytes | Resident memory size in bytes. | GAUGE |
| process_start_time_seconds | Start time of the process since unix epoch in seconds. | GAUGE |
| process_virtual_memory_bytes | Virtual memory size in bytes. | GAUGE |
| process_virtual_memory_max_bytes | Maximum amount of virtual memory available in bytes. | GAUGE |

