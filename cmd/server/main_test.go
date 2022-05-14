package main

import (
	createIncidentCommands "emergence/commandsAndQueries/CreateIncident"
	createTeamCommands "emergence/commandsAndQueries/CreateTeam"
	getCityAndAddressByExternalUserIdCommand "emergence/commandsAndQueries/GetCityAndAddressByExternalUserId"
	"emergence/integration"
	"emergence/integration/clients"
	"fmt"
	"github.com/baranius/godiator"
	"testing"
)

func TestAddingATeam(t *testing.T) {
	//create the mediator for the application and register the commands and queries
	postgresClient := clients.NewPostgresClient()
	opsGenieClient := clients.NewOpsGenieClient()
	repo := integration.NewRepository(postgresClient, opsGenieClient)
	g := godiator.GetInstance()

	g.Register(
		&createTeamCommands.CreateTeamCommand{},
		func() interface{} {
			return &createTeamCommands.CreateTeamCommandHandler{
				IRepository: repo,
			}
		})

	//register CreateIncidentCommandHandler
	g.Register(
		&createIncidentCommands.CreateIncidentCommand{},
		func() interface{} {
			return &createIncidentCommands.CreateIncidentCommandHandler{
				IRepository: repo,
			}
		})
	//register GetCityAndAddressCommandHandler
	g.Register(
		&getCityAndAddressByExternalUserIdCommand.GetCityAndAddressByExternalUserIdCommand{},
		func() interface{} {
			return &getCityAndAddressByExternalUserIdCommand.GetCityAndAddressByExternalUserIdCommandHandler{
				IRepository: repo,
			}
		})
	res, err := g.Send(&createTeamCommands.CreateTeamCommand{
		City:   "test",
		Emails: []string{"fdfd@gmail.com"},
	})
	if err != nil {
		fmt.Println(res)
	}
}

func TestClearTeamBlocking(t *testing.T) {
	//create the mediator for the application and register the commands and queries
	postgresClient := clients.NewPostgresClient()
	opsGenieClient := clients.NewOpsGenieClient()
	repo := integration.NewRepository(postgresClient, opsGenieClient)
	err := repo.SetTeamIsBlockedFalse("aboba")
	if err != nil {
		t.Error("error setting team to false")
	}
}
