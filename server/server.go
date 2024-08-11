package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/aofattaporn/go-cobra/configs"
	"github.com/aofattaporn/go-cobra/internal/middlewares"
	"github.com/aofattaporn/go-cobra/pkg/database"
	"github.com/aofattaporn/go-cobra/pkg/log"
	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start()
}

type fiberServer struct {
	cfg    configs.IConfig
	logger log.ILogger
	db     *sql.DB
	app    *fiber.App
}

func NewFiberServer(cfg configs.IConfig) (IServer, error) {

	logger, err := log.InitZapLogger(cfg.Log())
	if err != nil {
		return nil, fmt.Errorf("initializing logger error: %+v", err)
	}

	db, err := database.NewMysqlDatabase(cfg.Db(), logger)
	if err != nil {
		return nil, err
	}

	return &fiberServer{
		cfg:    cfg,
		logger: logger,
		db:     db,
		app: fiber.New(fiber.Config{
			AppName:               cfg.App().Name(),
			BodyLimit:             cfg.App().BodyLimit(),
			ReadTimeout:           cfg.App().ReadTimeout(),
			WriteTimeout:          cfg.App().WriteTimeout(),
			JSONEncoder:           json.Marshal,
			JSONDecoder:           json.Unmarshal,
			DisableStartupMessage: true,
			ErrorHandler:          middlewares.MappingError(logger),
		}),
	}, nil
}

func (s *fiberServer) Start() {
	logger := s.logger

	mid := InitMiddlewares(s)
	s.app.Use(mid.Cors())
	s.app.Use(mid.Logger())
	s.app.Use(recover.New(recover.Config{
		Next:              nil,
		EnableStackTrace:  true,
		StackTraceHandler: mid.Recover(),
	}))

	router := s.app.Group(s.cfg.App().ContextPath())
	modules := InitModule(router, s, mid)

	modules.HealthCheckModule()

	s.app.Use(mid.RouterNotFound())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		logger.Info("server is shutting down...")
		s.db.Close()
		_ = s.app.Shutdown()
	}()

	logger.Infof("server is starting on %d", s.cfg.App().Port())
	s.app.Listen(fmt.Sprintf(":%d", s.cfg.App().Port()))
}
