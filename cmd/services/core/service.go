package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/external/api"
	"github.com/khanghld27/kelvin-kart-challenge-api/cmd/services/common"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"

	"github.com/gin-gonic/gin"

	"github.com/kelseyhightower/envconfig"
)

type application struct {
	cfg    *configs.Config
	engine *gin.Engine
}

func (a *application) start() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	common.InitDBConnection(a.cfg.PostgreSQL)
	a.initJWTSession(a.cfg.JWTSecret)

	a.engine = api.Restful(a.cfg)

	go func() {
		if err := a.engine.Run(":" + a.cfg.HTTPServer.Port); err != nil {
			panic(err)
		}
	}()

	<-interrupt

	a.stopping()

	return nil
}

// stopping will stop running job or release resources the server was used
func (a *application) stopping() {
	logger.Debug("server stopped")
}

func newService() *application {
	s := &application{}
	s.loadConfig()

	logger.Init(s.cfg.LogLevel == string(logger.DebugLevel))
	logger.SetLevel(s.cfg.LogLevel)

	logger.Error("Some error happens")
	return s
}

func (a *application) loadConfig() {
	var cfg configs.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logrus.Fatal(err)
	}

	configs.SetConfig(&cfg)
	a.cfg = &cfg
}
