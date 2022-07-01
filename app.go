package main

import (
	"livecode-gorm-wmb/config"
)

func main() {
	config := config.NewConfigDB()
	// db := config.DbConn()
	defer config.DbClose()

	// migration.MigrateDb(db)

	// OAL 1
	// menuRepo := repository.NewMenuRepository(db)
	// command.MenuCreate(menuRepo)
	// command.MenuRead(menuRepo)
	// command.MenuUpdate(menuRepo)
	// command.MenuDelete(menuRepo)

	// SOAL 2
	// menuPriceRepo := repository.NewMenuPriceRepository(db)
	// command.MenuPriceCreate(menuPriceRepo)
	// command.MenuPriceRead(menuPriceRepo)
	// command.MenuPriceUpdate(menuPriceRepo)
	// command.MenuPriceDelete(menuPriceRepo)

	// SOAL 3
	// tableRepo := repository.NewTableRepository(db)
	// command.TableCreate(tableRepo)
	// command.TableRead(tableRepo)
	// command.TableUpdate(tableRepo)
	// command.TableDelete(tableRepo)

	// SOAL 4
	// menuPriceRepo := repository.NewMenuPriceRepository(db)
	// command.MenuPriceCreate(menuPriceRepo)
	// command.MenuPriceRead(menuPriceRepo)
	// command.MenuPriceUpdate(menuPriceRepo)
	// command.MenuPriceDelete(menuPriceRepo)

	// SOAL 5
	// menuPriceRepo := repository.NewMenuPriceRepository(db)
	// command.MenuPriceCreate(menuPriceRepo)
	// command.MenuPriceRead(menuPriceRepo)
	// command.MenuPriceUpdate(menuPriceRepo)
	// command.MenuPriceDelete(menuPriceRepo)

}

// DROP SCHEMA public CASCADE;
// CREATE SCHEMA public;
