package config

func GetMigrationDbConfig() DbConfig {
	const (
		defaultMigrationHost     = "localhost"
		defaultMigrationPort     = 5432
		defaultMigrationUser     = "postgres"
		defaultMigrationPassword = "password"
		defaultMigrationDbName   = "postgres"
	)
	return DbConfig{
		Host:     GetEnv("DB_HOST", defaultMigrationHost),
		Port:     GetEnvAsInt("DB_PORT", defaultMigrationPort),
		User:     GetEnv("DB_USER", defaultMigrationUser),
		Password: GetEnv("DB_PASSWORD", defaultMigrationPassword),
		DbName:   GetEnv("DB_NAME", defaultMigrationDbName),
	}
}

func GetAppDbConfig() DbConfig {
	const (
		defaultAppHost     = "localhost"
		defaultAppPort     = 5432
		defaultAppUser     = "postgres"
		defaultAppPassword = "password"
		defaultAppDbName   = "postgres"
	)
	return DbConfig{
		Host:     GetEnv("DB_HOST", defaultAppHost),
		Port:     GetEnvAsInt("DB_PORT", defaultAppPort),
		User:     GetEnv("APP_USER", defaultAppUser),
		Password: GetEnv("APP_PASSWORD", defaultAppPassword),
		DbName:   GetEnv("DB_NAME", defaultAppDbName),
	}
}
