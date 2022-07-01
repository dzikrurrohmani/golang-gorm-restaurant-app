package main

import (
	"livecode-gorm-wmb/config"
	"livecode-gorm-wmb/migration"
)

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer config.DbClose()

	migration.MigrateDb(db)

	// SOAL 1
	// custRepo := repository.NewCustomerRepository(db)

}

// DROP SCHEMA public CASCADE;
// CREATE SCHEMA public;
