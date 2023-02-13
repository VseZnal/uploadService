package config

import "uploadService/libs/utils"

var (
	uploadServiceHostEnvName = "UPLOAD_SERVICE_HOST"
	uploadServicePortEnvName = "UPLOAD_SERVICE_PORT"
	pgConnStringEnvName      = "PG_CONNECTION_STRING"
)

type Config struct {
	HostUpload   string
	PortUpload   string
	PgConnString string
}

func UploadConfig() *Config {
	return &Config{
		HostUpload:   utils.TrimEnv(uploadServiceHostEnvName),
		PortUpload:   utils.TrimEnv(uploadServicePortEnvName),
		PgConnString: utils.TrimEnv(pgConnStringEnvName),
	}
}
