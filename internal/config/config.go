package config

import (
	"os"
	"strconv"
)

const (
	DefaultMongoHost     = "localhost"
	DefaultMongoPort     = 27017
	DefaultMongoUsername = "admin"
	DefaultMongoPassword = "password"
	DefaultBackupDir     = "/tmp/micro-database-backup/mongodb"
	DefaultDatabaseName  = "mydb"
)

var (
	MongoHost     = getEnv("MONGODB_HOST", DefaultMongoHost)
	MongoPort     = getEnvAsInt("MONGODB_PORT", DefaultMongoPort)
	MongoUsername = getEnv("MONGODB_USERNAME", DefaultMongoUsername)
	MongoPassword = getEnv("MONGODB_PASSWORD", DefaultMongoPassword)
	BackupDir     = getEnv("BACKUP_DIR", DefaultBackupDir)
	DatabaseName  = getEnv("DATABASE_NAME", DefaultDatabaseName)
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
