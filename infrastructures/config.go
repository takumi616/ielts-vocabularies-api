package infrastructures

import (
	"github.com/caarlos0/env/v11"
)

type environmentVariables struct {
	AppPort  string `env:"APP_CONTAINER_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	PassWord string `env:"POSTGRES_PASSWORD"`
	Sslmode  string `env:"POSTGRES_SSLMODE"`
}

type Config struct {
	AppPort  string
	PgConfig *PgConfig
}

type PgConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	PassWord string
	Sslmode  string
}

// Read environment variables and set them to golang struct
func NewConfig() (*Config, error) {
	environmentVariables := &environmentVariables{}
	//Set environment variables to environmentVariables.
	if err := env.Parse(environmentVariables); err != nil {
		return nil, err
	}

	pgConfig := &PgConfig{
		Host:     environmentVariables.Host,
		Port:     environmentVariables.Port,
		DbName:   environmentVariables.DbName,
		User:     environmentVariables.User,
		PassWord: environmentVariables.PassWord,
		Sslmode:  environmentVariables.Sslmode,
	}

	config := &Config{
		AppPort:  environmentVariables.AppPort,
		PgConfig: pgConfig,
	}

	return config, nil
}
