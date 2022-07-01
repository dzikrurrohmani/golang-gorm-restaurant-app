package main

import "livecode-gorm-wmb/config"

func main() {
	config := config.NewConfigDB()
	// db := config.DbConn()
	config.DbConn() // kebutuhan migration
	defer config.DbClose()
}
