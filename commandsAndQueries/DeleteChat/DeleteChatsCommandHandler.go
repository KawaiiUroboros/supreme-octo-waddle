package commandsAndQueries

// Handle handler for delete chats by user ids using IRepository
func (h *DeleteChatsCommandHandler) Handle(request *DeleteChatsCommand) (*DeleteChatsResult, error) {
	result, err := h.IRepository.DeleteChats(request.ExternalUserIds)
	if err != nil {
		return &DeleteChatsResult{}, err
	}
	return &DeleteChatsResult{
		DeletedChatIds: result,
	}, nil
}
