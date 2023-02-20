package config

import "uploadService/libs/utils"

var (
	uploadClientHostEnvName = "UPLOAD_SERVICE_HOST"
	uploadClientPortEnvName = "UPLOAD_SERVICE_PORT"
)

type Config struct {
	HostUpload string
	PortUpload string
}

func UploadClientConfig() *Config {
	return &Config{
		HostUpload: utils.TrimEnv(uploadClientHostEnvName),
		PortUpload: utils.TrimEnv(uploadClientPortEnvName),
	}
}
