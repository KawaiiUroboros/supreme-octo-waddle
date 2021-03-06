// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"emergence/integration"
	"emergence/integration/clients"
)

// Injectors from wire.go:

// InitializeDependency initialize all structs and interfaces
func InitializeDependency() (integration.IRepository, error) {
	postgresClient := clients.NewPostgresClient()
	opsGenieClient := clients.NewOpsGenieClient()
	repository := integration.NewRepository(postgresClient, opsGenieClient)
	return repository, nil
}
