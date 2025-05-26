package configs

import (
	"fmt"
	"sync"
	"time"
)

var (
	once      sync.Once
	singleton *Config
)

// SetConfig to set configuration of service.
func SetConfig(cfg *Config) *Config {
	once.Do(func() {
		singleton = cfg
	})
	return singleton
}

// GetConfig gets the instance of singleton
func GetConfig() *Config {
	return singleton
}

// PostgreSQL config for PostgreSQL server
type PostgreSQL struct {
	User            string `envconfig:"DB_USER"`
	Password        string `envconfig:"DB_PASSWORD"`
	Host            string `envconfig:"DB_HOST"`
	Port            string `envconfig:"DB_PORT"`
	Database        string `envconfig:"DB_DATABASE"`
	SSLMode         string `envconfig:"DB_SSL_MODE"`
	MaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime int    `envconfig:"DB_CONN_MAX_LIFETIME"`
	IsEnabledLog    bool   `envconfig:"DB_IS_ENABLED_LOG"`
}

// Conn return connection string
func (p *PostgreSQL) Conn() string {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.Database,
	)
	return dbUri
}

// ServerCfg server addresses
type ServerCfg struct {
	Port    string        `envconfig:"PORT" default:"8080"`
	Timeout time.Duration `envconfig:"TIMEOUT" default:"3s"`
}

// Config is APP config information
type Config struct {
	Env              string `envconfig:"ENV" default:"local"`
	HTTPServer       ServerCfg
	PostgreSQL       PostgreSQL
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info"`
	JWTSecret        string `envconfig:"JWT_SECRET" default:"secret"`
	EnabledProfiling bool   `envconfig:"ENABLE_PROFILING" default:"false"`
}
