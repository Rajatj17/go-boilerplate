package config

type DbConfig struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
	DbDialect  string
	DbSslMode  bool
}
