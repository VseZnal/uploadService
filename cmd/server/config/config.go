package config

import "uploadService/libs/utils"

var (
	uploadServerPortEnvName = "UPLOAD_SERVICE_PORT"
	staticDirName           = "STATIC_DIR"
	pgConnStringEnvName     = "PG_CONNECTION_STRING"
)

type Config struct {
	PortServer   string
	Static       string
	PgConnString string
}

func UploadServerConfig() *Config {
	return &Config{
		PortServer:   utils.TrimEnv(uploadServerPortEnvName),
		Static:       utils.TrimEnv(staticDirName),
		PgConnString: utils.TrimEnv(pgConnStringEnvName),
	}
}
