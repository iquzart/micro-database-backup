package main

import (
	"micro-database-backup/internal/config"
	"micro-database-backup/internal/database/mongodb"
	"micro-database-backup/internal/logger"
	"path/filepath"
	"time"
)

func main() {
	log := logger.InitLogger()

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(config.BackupDir, config.DatabaseName+"_"+timestamp)

	log.Info("Identified backup path", "backup_path", backupPath)

	// Create MongoDB database instance
	mongoDB := mongodb.NewMongoDB()

	// Create backup directory
	if err := mongoDB.CreateBackupDirectory(backupPath); err != nil {
		handleError(log, "Error creating backup directory", err)
		return
	}

	log.Info("Backup directory created successfully", "backup_path", backupPath)

	// Perform database backup
	if err := mongoDB.BackupDatabase(config.DatabaseName, backupPath); err != nil {
		handleError(log, "Error performing database backup", err)
		return
	}

	log.Info("Database backup completed successfully", "backup_path", backupPath)
}

func handleError(log *logger.Logger, message string, err error) {
	log.Error(message, "error", err)
	// Implement further error handling logic, such as sending email notifications
}
