package commandsAndQueries

func (h *CreateIncidentCommandHandler) Handle(c *CreateIncidentCommand) (*CreateIncidentCommandResult, error) {
	err := h.IRepository.CreateIncident(&c.Incident)
	if err != nil {
		return nil, err
	}
	return &CreateIncidentCommandResult{Success: true}, nil
}
