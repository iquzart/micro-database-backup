package database

// Database interface defines methods for interacting with different database systems
type Database interface {
	// CreateBackupDirectory creates a backup directory for the database
	CreateBackupDirectory(backupPath string) error

	// CreateBackup performs a backup of the database
	CreateBackup(backupPath string) error

	// CleanupBackups deletes old backup files from the specified directory
	// CleanupBackups(backupDir, databaseName string) error
}
