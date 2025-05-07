package main

import (
	"bondly/migrations"
	"bondly/utils"
)

func main() {
	// Initialize logger
	utils.InitLogger()
	appLogger := utils.GetLogger()

	// Log migration start
	appLogger.Info("Starting SQL migration...")

	// Initialize database connection
	_, err := utils.InitDatabase() // No need for the 'cfg' variable
	if err != nil {
		appLogger.Fatalf("Error initializing the database: %v", err)
	}

	// Run migrations
	if err := migrations.MigrateSQL(); err != nil { // Call the function without arguments
		appLogger.Fatalf("Error running migrations: %v", err)
	}

	appLogger.Info("Migrations completed successfully!")
}
