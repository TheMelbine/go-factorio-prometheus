package meters

func Setup(manager *Manager) {
	PlayerMeters(manager)
	PlanetsMeters(manager)
	ForcesMeters(manager)
	LogisticsMeters(manager)
}
