// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/ameteiko/mindnet/domain/entity"
	"github.com/ameteiko/mindnet/domain/service"
	"github.com/ameteiko/mindnet/domain/value"
	"github.com/ameteiko/mindnet/domain/value/sanitiser"
	"github.com/ameteiko/mindnet/domain/value/validator"
	"github.com/ameteiko/mindnet/internal/app"
	neo4j2 "github.com/ameteiko/mindnet/internal/platform/neo4j"
	"github.com/ameteiko/mindnet/internal/platform/services"
)

// Injectors from wire.go:

func provideValueSanitiser() sanitiser.Sanitiser {
	sanitiserSanitiser := sanitiser.Sanitiser{}
	return sanitiserSanitiser
}

func provideValueValidator() validator.Validator {
	validatorValidator := validator.Validator{}
	return validatorValidator
}

func provideIDGenerator() services.IDGenerator {
	idGenerator := services.IDGenerator{}
	return idGenerator
}

func provideNodeRepository(dbSession neo4j.Session) *neo4j2.DB {
	db := neo4j2.NewDB(dbSession)
	return db
}

func provideNodeService() service.Node {
	node := service.Node{}
	return node
}

func provideValueFactory() value.Factory {
	sanitiserSanitiser := provideValueSanitiser()
	validatorValidator := provideValueValidator()
	factory := value.NewFactory(sanitiserSanitiser, validatorValidator)
	return factory
}

func provideEntityFactory() entity.Factory {
	factory := provideValueFactory()
	idGenerator := provideIDGenerator()
	entityFactory := entity.NewFactory(factory, idGenerator)
	return entityFactory
}

func ProvideApp(dbSession neo4j.Session) app.App {
	factory := provideEntityFactory()
	db := provideNodeRepository(dbSession)
	node := provideNodeService()
	appApp := app.New(factory, db, node)
	return appApp
}
