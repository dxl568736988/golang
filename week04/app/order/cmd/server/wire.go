// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"homework/week04/app/order/internal/biz"
	"homework/week04/app/order/internal/conf"
	"homework/week04/app/order/internal/data"
	"homework/week04/app/order/internal/server"
	"homework/week04/app/order/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
