package db

import (
	"fmt"
	"github.com/Laky-64/TestFlightTrackBot/internal/config"
	"github.com/Laky-64/TestFlightTrackBot/internal/db/models"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

type DB struct {
	AppStore      *appStore
	LinkStore     *linkStore
	PendingStore  *pendingStore
	ChatLinkStore *chatLinkStore
}

func NewDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Rome",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	goose.SetLogger(goose.NopLogger())
	if err = goose.Up(sqlDB, "internal/db/migrations"); err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&models.Chat{},
		&models.App{},
		&models.Link{},
		&models.ChatLink{},
		&models.PendingTrack{},
		&models.PremiumUser{},
		&models.TranslatorLanguage{},
	); err != nil {
		return nil, err
	}

	pending := &pendingStore{
		db:              db,
		pendingTracking: sync.Map{},
	}

	if err = pending.LoadFromDB(); err != nil {
		return nil, err
	}

	return &DB{
		AppStore:      &appStore{db: db},
		LinkStore:     &linkStore{db: db, cfg: cfg},
		PendingStore:  pending,
		ChatLinkStore: &chatLinkStore{db: db, cfg: cfg},
	}, nil
}
