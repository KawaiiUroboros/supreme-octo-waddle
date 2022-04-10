package main

import (
	"emergence/app/server"
	"emergence/cmd/migrations"
	createChatCommands "emergence/commandsAndQueries/CreateChat"
	deleteChatCommands "emergence/commandsAndQueries/DeleteChat"
	"github.com/baranius/godiator"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	panicHandler "github.com/kazegusuri/grpc-panic-handler"
	"github.com/srvc/appctx"
	"log"
	os "os"
	"time"
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
		&createChatCommands.CreateChatCommand{},
		func() interface{} {
			return &createChatCommands.CreateChatCommandHandler{
				IRepository: repo,
			}
		})

	//register DeleteChatCommand
	g.Register(
		&deleteChatCommands.DeleteChatsCommand{},
		func() interface{} {
			return &deleteChatCommands.DeleteChatsCommandHandler{
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
			server.NewNotyfierServiceServer(repo),
		),
	)

	//create a background task to send notifications to chats that have exceeded the interval evry minute
	go func() {
		for {
			log.Println("Sending notifications to chats")
			err := repo.NotifyChats()
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Minute)
		}
	}()

	return s.Serve(ctx)
}
