package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type appEnv string

const (
	AppEnvDevelopment appEnv = "development"
	AppEnvProduction  appEnv = "production"
)

type EnvVars struct {
	Port   string `env:"PORT,required"`
	AppEnv appEnv `env:"APP_ENV,required"`
}

var Env *EnvVars

func init() {
	Env = &EnvVars{}
	if err := env.Parse(Env); err != nil {
		panic(fmt.Sprintf("❌ Invalid Environment Variable: %s", err))
	}
}
