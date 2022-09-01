package migration

import (
	"github.com/felixlambertv/online-attendance/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(model.User{})
	if err != nil {
		logrus.Error("Fail to migrate table")
	}
}
