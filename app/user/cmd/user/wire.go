//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/app/user/internal/conf"
	"github.com/qqz/Happy-Stream/app/user/internal/data"
	"github.com/qqz/Happy-Stream/app/user/internal/server"
	"github.com/qqz/Happy-Stream/app/user/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
