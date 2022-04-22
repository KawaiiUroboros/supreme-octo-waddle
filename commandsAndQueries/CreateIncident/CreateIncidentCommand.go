package commandsAndQueries

import (
	"emergence/commandsAndQueries/models"
	"emergence/integration"
)

type (
	CreateIncidentCommand struct {
		Incident models.Incident
	}
	CreateIncidentCommandResult struct {
		Success bool
	}
	CreateIncidentCommandHandler struct {
		IRepository integration.IRepository
	}
)
