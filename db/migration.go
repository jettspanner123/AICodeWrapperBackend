package db

import (
	"log"

	activityModel "github.com/jettspanner123/AICodeWrapperBackend/models/activity"
)

func MigrateModels() {
	err := Database.AutoMigrate(
		&activityModel.APILog{},
	)

	if err != nil {
		log.Fatal("Migration Failed! ❌", err)
	}

	log.Println("Database Migration Successful! ✅")
}
