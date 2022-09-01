package main

import (
	"github.com/felixlambertv/online-attendance/config"
	"github.com/felixlambertv/online-attendance/migration"
)

func main() {
	db := config.SetupDb()
	migration.Migrate(db)
	config.SetupServer(db)
}
