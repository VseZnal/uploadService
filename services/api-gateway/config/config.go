package config

import "uploadService/libs/utils"

var (
	gatewayHostEnvName       = "API_HOST"
	gatewayPortEnvName       = "API_PORT"
	uploadServiceHostEnvName = "UPLOAD_SERVICE_HOST"
	uploadServicePortEnvName = "UPLOAD_SERVICE_PORT"
	cors                     = "cors"
)

type Config struct {
	HostGateway string
	PortGateway string
	HostUpload  string
	PortUpload  string
	Cors        string
}

func GatewayConfig() *Config {
	return &Config{
		HostGateway: utils.TrimEnv(gatewayHostEnvName),
		PortGateway: utils.TrimEnv(gatewayPortEnvName),
		HostUpload:  utils.TrimEnv(uploadServiceHostEnvName),
		PortUpload:  utils.TrimEnv(uploadServicePortEnvName),
		Cors:        utils.TrimEnv(cors),
	}
}
