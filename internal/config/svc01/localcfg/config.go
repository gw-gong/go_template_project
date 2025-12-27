package localcfg

import (
	"github.com/gw-gong/go-template-project/internal/config/types"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz01"
	"github.com/gw-gong/go-template-project/internal/pkg/biz/biz02"

	"github.com/gw-gong/gwkit-go/gin/middleware"
	"github.com/gw-gong/gwkit-go/hotcfg"
	"github.com/gw-gong/gwkit-go/log"
	"github.com/gw-gong/gwkit-go/setting"
)

type Config struct {
	hotcfg.BaseConfigCapable
	Env          setting.Env                    `yaml:"env" mapstructure:"env"`
	HttpServer   *types.HttpServer              `yaml:"http_server" mapstructure:"http_server"`
	ConsulNetCfg *hotcfg.ConsulConfigOption     `yaml:"consul_net_cfg" mapstructure:"consul_net_cfg"`
	Logger       *log.LoggerConfig              `yaml:"logger" mapstructure:"logger"`
	LogHttpInfo  *middleware.LogHttpInfoOptions `yaml:"log_http_info" mapstructure:"log_http_info"`
	Biz01        *biz01.Biz01Options            `yaml:"biz01" mapstructure:"biz01"`
	Biz02        *biz02.Biz02Options            `yaml:"biz02" mapstructure:"biz02"`
}

func (c *Config) LoadConfig() {
	if err := c.Unmarshal(&c); err != nil {
		log.Error("unmarshal config failed", log.Err(err))
		return
	}

	log.Info("LoadConfig", log.Any("config", c))
}

func NewConfig(localConfigOption *hotcfg.LocalConfigOption) (config *Config, err error) {
	config = &Config{}
	config.BaseConfigCapable, err = hotcfg.NewLocalBaseConfigCapable(localConfigOption)
	return config, err
}
