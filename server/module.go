package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
	"github.com/goproject/internal/middlewares"
)

type IModuleFactory interface {
	HealthCheckModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *fiberServer
	mid middlewares.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *fiberServer, mid middlewares.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *fiberServer) middlewares.IMiddlewaresHandler {
	return middlewares.MiddlewaresHandler(s.cfg, s.logger)
}

func (m *moduleFactory) HealthCheckModule() {
	h := handlers.HealthCheckHandler(m.s.cfg.App())

	fmt.Println("Checking 2")
	m.r.Get(constants.ROUTE().HEALTHCHECK, h.HeathCheckHandler)
}
