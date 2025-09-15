package service01

import (
	"github.com/gw-gong/template_project/internal/pkg/components/component01"
	"github.com/gw-gong/template_project/internal/pkg/components/component02"
	"github.com/gw-gong/gwkit-go/log"
)

type Config struct {
	LogConfig            *log.LoggerConfig
	Component01erOptions component01.Component01erOptions
	Component02erOptions component02.Component02erOptions
}

func InitConfig() (*Config, error) {
	return nil, nil
}
