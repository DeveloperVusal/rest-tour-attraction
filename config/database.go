package config

import (
	"attrtour/core"
)

type DatabaseConfig struct{}

func (dbcfg *DatabaseConfig) ConfigLoad() map[string]string {
	env := core.Helpers{}

	return map[string]string{
		"host":     env.Env("DB_HOST"),
		"port":     env.Env("DB_PORT"),
		"username": env.Env("DB_USERNAME"),
		"password": env.Env("DB_PASSWORD"),
		"dbname":   env.Env("DB_DATABASE"),
	}
}
