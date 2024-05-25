package mongodb

import (
	"fmt"
	"micro-database-backup/internal/config"
	"os"
	"os/exec"
	"path/filepath"
)

type MongoDB struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMongoDB() *MongoDB {
	return &MongoDB{
		Host:     config.MongoHost,
		Port:     config.MongoPort,
		Username: config.MongoUsername,
		Password: config.MongoPassword,
	}
}

func (db *MongoDB) CreateBackupDirectory(backupPath string) error {
	return os.MkdirAll(backupPath, os.ModePerm)
}

func (db *MongoDB) BackupDatabase(databaseName, backupPath string) error {
	backupDir := filepath.Join(backupPath, databaseName)
	cmd := exec.Command("mongodump",
		"--host", db.Host,
		"--port", fmt.Sprint(db.Port),
		"--username", db.Username,
		"--password", db.Password,
		"--db", databaseName,
		"--out", backupDir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
