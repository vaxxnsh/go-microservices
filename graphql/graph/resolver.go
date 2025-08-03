package graph

import (
	"github.com/vaxxnsh/go-microservices/graphql/server"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Server *server.Server
}
