package commandsAndQueries

import "emergence/integration"

//command, result, command handler structs for deleting chats by string external user ids
type (
	DeleteChatsCommand struct {
		ExternalUserIds *[]string
	}
	DeleteChatsResult struct {
		DeletedChatIds *[]string
	}
	DeleteChatsCommandHandler struct {
		IRepository integration.IRepository
	}
)
