package commandsAndQueries

import "emergence/integration"

type (
	GetCityAndAddressByExternalUserIdCommand struct {
		ExternalUserId string `json:"externalUserId"`
	}

	GetCityAndAddressByExternalUserIdCommandResult struct {
		// City -
		City string `json:"city"`
		// Address -
		Address string `json:"address"`
	}

	GetCityAndAddressByExternalUserIdCommandHandler struct {
		IRepository integration.IRepository
	}
)
