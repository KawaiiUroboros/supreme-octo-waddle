package integration

import (
	"emergence/commandsAndQueries/models"
	"emergence/integration/clients"
)

//IRepository interface for creating Slack channels and storing them in the database
type IRepository interface {
	CreateTeam(emails []string, city string) (teamId string, error error)
	CreateIncident(incident *models.Incident) (error error)
	GetCityAndAddressByExternalUserId(externalUserId string) (city string, address string, err error)
}

// Repository implementation of the IRepository interface for the database and OpsGenie
type Repository struct {
	db       clients.IPostgresClient
	opsGenie clients.IOpsGenieClient
}

//NewRepository creates a new Repository
func NewRepository(db clients.IPostgresClient, opsGenie clients.IOpsGenieClient) *Repository {
	return &Repository{
		db:       db,
		opsGenie: opsGenie,
	}
}

//CreateTeam creates a new team in OpsGenie and then add its id and city to the database
func (r *Repository) CreateTeam(emails []string, city string) (teamId string, error error) {
	teamId, err := r.opsGenie.CreateTeam(emails)
	if err != nil {
		return "", err
	}
	err = r.db.AddTeam(teamId, city)
	if err != nil {
		return "", err
	}
	return teamId, nil
}

//CreateIncident get the team by city from the database and create an incident in OpsGenie
func (r *Repository) CreateIncident(incident *models.Incident) error {
	team, err := r.db.GetTeamByCity(incident.City)
	if err != nil {
		return err
	}
	err = r.opsGenie.CreateIncident(incident, team.Id)
	if err != nil {
		return err
	}
	return nil
}

//GetCityAndAddressByExternalUserId gets the city and address by external user id from the database
func (r *Repository) GetCityAndAddressByExternalUserId(externalUserId string) (city string, address string, err error) {
	return r.db.GetCityAndAddressByExternalUserId(externalUserId)
}
