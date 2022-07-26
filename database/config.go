package database

import (
	"fmt"

	env "github.com/caitlinelfring/go-env-default"
)

type dbConfig struct {
	host string
	port int
	user string
	password string
	database string
}

func (c *dbConfig) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", 
		c.host, 
		c.user, 
		c.password, 
		c.database,
		c.port)
}

func configFromEnv() *dbConfig {
	config := new(dbConfig)

	config.host = env.GetDefault("POSTGRES_HOST", "postgres")
	config.port = env.GetIntDefault("POSTGRES_PORT", 5432)
	config.user = env.GetDefault("POSTGRES_USER", "postgres")
	config.password = env.GetDefault("POSTGRES_PASSWORD", "postgres")
	config.database = env.GetDefault("POSTGRES_DATABASE", "postgres")

	return config
}