package commandsAndQueries

// Handle create team in db and ops genie
func (h *CreateTeamCommandHandler) Handle(command *CreateTeamCommand) (*CreateTeamResult, error) {

	_, err := h.IRepository.CreateTeam(command.Emails, command.City)
	if err != nil {
		return nil, err
	}

	return &CreateTeamResult{
		Success: true,
	}, nil
}
