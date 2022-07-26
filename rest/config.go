package rest

import "github.com/caitlinelfring/go-env-default"

type restConfig struct {
	port int
}

const UserServiceDefaultPort = 3333

func configFromEnv() *restConfig {
	config := new(restConfig)
	config.port = env.GetIntDefault("HTTP_PORT", UserServiceDefaultPort)
	return config
}