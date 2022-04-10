package commandsAndQueries

import "emergence/integration"

type (
	CreateChatCommand struct {
		ExternalUserIdToInterval *map[string]int
	}

	CreateChatResult struct {
		ExternalUserIdToChatUrl *map[string]string
	}

	CreateChatCommandHandler struct {
		IRepository integration.IRepository
	}
)
