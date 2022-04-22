package commandsAndQueries

// Handle get city and address by external user id
func (h *GetCityAndAddressByExternalUserIdCommandHandler) Handle(command *GetCityAndAddressByExternalUserIdCommand) (*GetCityAndAddressByExternalUserIdCommandResult, error) {
	address, city, err := h.IRepository.GetCityAndAddressByExternalUserId(command.ExternalUserId)
	if err != nil {
		return nil, err
	}
	return &GetCityAndAddressByExternalUserIdCommandResult{
		Address: address,
		City:    city,
	}, nil
}
