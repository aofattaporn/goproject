package server

import (
	"github.com/aofattaporn/go-cobra/internal/handlers"
	"github.com/aofattaporn/go-cobra/internal/middlewares"
	"github.com/gofiber/fiber/v2"
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
	m.r.Get(constants.ROUTE().HEALTHCHECK, h.HeathCheckHandler)
}
