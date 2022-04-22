package commandsAndQueries

import "emergence/integration"

type (
	CreateTeamCommand struct {
		City   string
		Emails []string
	}

	CreateTeamResult struct {
		Success bool
	}

	CreateTeamCommandHandler struct {
		IRepository integration.IRepository
	}
)
