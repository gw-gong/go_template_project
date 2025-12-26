package types

type HttpServer struct {
	Port int         `yaml:"port" mapstructure:"port"`
	Cors *CorsConfig `yaml:"cors" mapstructure:"cors"`
}

type CorsConfig struct {
	AllowOrigins     []string `yaml:"allow_origins" mapstructure:"allow_origins"`
	AllowCredentials bool     `yaml:"allow_credentials" mapstructure:"allow_credentials"`
}
