package gormer

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
)

type DBAdapter interface {
	Connect(connStr string, config gorm.Config) error
	Close()
	Begin() DBAdapter
	RollbackUselessCommitted()
	Gormer() *gorm.DB
	DB() *sql.DB
}

type adapter struct {
	gormer      *gorm.DB
	isCommitted bool
}

func NewDB() DBAdapter {
	return &adapter{}
}

func (db *adapter) Connect(connStr string, config gorm.Config) error {
	var (
		err    error
		gormer *gorm.DB
	)
	once.Do(func() {
		gormer, err = gorm.Open(postgres.Open(connStr), &config)
		if err != nil {
			return
		}
	})

	db.gormer = gormer
	return nil
}

func (db *adapter) Close() {
	_ = db.DB().Close()
}

// Begin starts a DB transaction.
func (db *adapter) Begin() DBAdapter {
	tx := db.gormer.Begin()

	return &adapter{
		gormer:      tx,
		isCommitted: false,
	}
}

// RollbackUselessCommitted rollbacks useless DB transaction committed.
func (db *adapter) RollbackUselessCommitted() {
	if !db.isCommitted {
		db.gormer.Rollback()
	}
}

// Commit commits a DB transaction.
func (db *adapter) Commit() {
	if !db.isCommitted {
		db.gormer.Commit()
		db.isCommitted = true
	}
}

// Gormer returns an instance of gorm.DB.
func (db *adapter) Gormer() *gorm.DB {
	return db.gormer
}

// DB returns an instance of sql.DB.
func (db *adapter) DB() *sql.DB {
	database, _ := db.gormer.DB()
	return database
}
