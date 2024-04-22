package databasencryptionscheduler

import (
	"time"

	"github.com/go-co-op/gocron"

	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
	"github.com/AmitKarnam/KeyCloak/internal/utils/scheduler"
)

type databasencryptionscheduler struct {
}

var _ scheduler.Scheduler = &databasencryptionscheduler{}

func Init() *databasencryptionscheduler {
	return &databasencryptionscheduler{}
}

// To Generate a new master key and store it in the mentioned file path.
var task = func() {
	// TODO : pre checks in place shoud be configured to check if the file is available, If file not available then exit.
	// If pre checks pass then , generate a new master key and store it in the file, store the old key in a .bak file.
	// Decrypt the data in the database using the old key, encrypt using the new key and then store the encrypte data back in the database.
}

func (m *databasencryptionscheduler) StartScheduler() error {

	zapLogger.KeyCloaklogger.Infof("Starting Master Key Scheduler")

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Cron(scheduler.MidNightSced).Do(task)
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("Error while running master key scheduler: %v", err)
		return err
	}

	s.StartAsync()
	return nil
}
