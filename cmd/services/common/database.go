package common

import (
	"time"

	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/gormer"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDBConnection(config configs.PostgreSQL) gormer.DBAdapter {
	logMode := logger.Default.LogMode(logger.Error)

	if config.IsEnabledLog {
		logMode = logger.Default.LogMode(logger.Info)
	}

	db := gormer.NewDB()

	if err := db.Connect(config.Conn(), gorm.Config{
		Logger: logMode,
	}); err != nil {
		logrus.Fatal("Creating connection to DB got error:", err)
	}

	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)

	return db
}
