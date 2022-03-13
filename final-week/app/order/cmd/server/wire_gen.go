// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"homework/week04/app/order/internal/biz"
	"homework/week04/app/order/internal/conf"
	"homework/week04/app/order/internal/data"
	"homework/week04/app/order/internal/server"
	"homework/week04/app/order/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUseCase := biz.NewOrderUseCase(orderRepo, logger)
	orderService := service.NewOrderService(orderUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, orderService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
