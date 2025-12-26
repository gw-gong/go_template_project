package common_config

import (
	"fmt"
	"sync"

	"github.com/gw-gong/go-template-project/internal/pkg/components/component01"
	"github.com/gw-gong/go-template-project/internal/pkg/components/component02"

	"github.com/gw-gong/gwkit-go/hot_cfg"
	"github.com/gw-gong/gwkit-go/log"
)

var (
	Cfg  *Config
	once sync.Once
)

type Config struct {
	*hot_cfg.BaseConfig
	Component01Options component01.Component01Options `yaml:"component01_options" mapstructure:"component01_options"`
	Component02Options component02.Component02Options `yaml:"component02_options" mapstructure:"component02_options"`
}

const (
	RootPath             = "../../"
	commLocalCfgPath     = RootPath + "config/common/local_config"
	commLocalCfgFileName = "config.yaml"
	commLocalCfgFileType = "yaml"
)

func InitConfig() error {
	var err error
	once.Do(func() {
		Cfg = &Config{}
		Cfg.BaseConfig, err = hot_cfg.NewBaseConfig(
			hot_cfg.WithLocalConfig(commLocalCfgPath, commLocalCfgFileName, commLocalCfgFileType),
		)
		if err != nil {
			err = fmt.Errorf("init base config failed: %w", err)
		}
	})
	return err
}

func (c *Config) LoadConfig() {

	if err := c.BaseConfig.Viper.Unmarshal(&c); err != nil {
		log.Error("unmarshal config failed", log.Err(err))
		return
	}

	log.Info("LoadConfig", log.Any("config", c))
}
