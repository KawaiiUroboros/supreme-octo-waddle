//go:build wireinject
// +build wireinject

package main

import (
	"emergence/integration"
	"emergence/integration/clients"
	"github.com/google/wire"
)

// InitializeDependency initialize all structs and interfaces
func InitializeDependency() (integration.IRepository, error) {
	wire.Build(
		clients.NewOpsGenieClient,
		clients.NewPostgresClient,
		integration.NewRepository,
		wire.Bind(new(clients.IOpsGenieClient), new(clients.OpsGenieClient)),
		wire.Bind(new(clients.IPostgresClient), new(clients.PostgresClient)),
		wire.Bind(new(integration.IRepository), new(integration.Repository)),
	)
	return &integration.Repository{}, nil
}
