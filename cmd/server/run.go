package main

import (
	"emergence/app/server"
	"emergence/cmd/migrations"
	createIncidentCommands "emergence/commandsAndQueries/CreateIncident"
	createTeamCommands "emergence/commandsAndQueries/CreateTeam"
	getCityAndAddressByExternalUserIdCommand "emergence/commandsAndQueries/GetCityAndAddressByExternalUserId"
	"github.com/baranius/godiator"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	panicHandler "github.com/kazegusuri/grpc-panic-handler"
	"github.com/srvc/appctx"
	"log"
	os "os"
)

func run() error {

	// Application context
	ctx := appctx.Global()
	//auto-migrations for the database
	errorFromMigrations := migrations.Compose()

	if errorFromMigrations != nil {
		log.Panic(errorFromMigrations)
	}
	repo, err := InitializeDependency()
	if err != nil {
		log.Panic(err)
	}

	//create the mediator for the application and register the commands and queries
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	s := grapiserver.New(
		grapiserver.WithGrpcAddr("tcp", ":9000"),
		grapiserver.WithGatewayAddr("tcp", ":"+port),
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcServerUnaryInterceptors(panicHandler.UnaryPanicHandler),
		grapiserver.WithServers(
			server.NewHospitalServiceServer(repo),
		),
	)
	return s.Serve(ctx)
}
