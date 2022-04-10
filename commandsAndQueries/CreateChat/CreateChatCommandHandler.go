package commandsAndQueries

import (
	"emergence/commandsAndQueries/models"
)

func (h *CreateChatCommandHandler) Handle(request *CreateChatCommand) (*CreateChatResult, error) {

	//create a slice of newChatOptions from ExternalUserIdToInterval
	newChatOptions := make([]models.NewChatOptions, 0)
	for key, value := range *request.ExternalUserIdToInterval {
		newChatOptions = append(newChatOptions, models.NewChatOptions{
			ExternalUserId: key,
			Interval:       value,
		})
	}

	//create Slack channels for each user and interval by using CreateChannelByUserIdAndInterval
	userIdToChannelId, err := h.IRepository.CreateChannelsByUserIdsAndIntervals(&newChatOptions)

	//if error return empty result and error
	if err != nil {
		return &CreateChatResult{}, err
	}
	return &CreateChatResult{
		ExternalUserIdToChatUrl: userIdToChannelId,
	}, err
}
