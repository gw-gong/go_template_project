//go:build wireinject
// +build wireinject

package components

import (
	"github.com/google/wire"
	"github.com/gw-gong/template_project/internal/pkg/components/component01"
	"github.com/gw-gong/template_project/internal/pkg/components/component02"
)

type Component struct {
	Component01 component01.Component01er
	Component02 component02.Component02er
}

// wire不支持输入相同类型的参数，所以套一层struct
type ComponentConfig struct {
	Component01Params struct {
		Field01 string
		Field02 string
	}
	Component02Params struct {
		Field01 string
		Field02 string
	}
}

// 初始化组件
func InitComponent(config ComponentConfig) (*Component, error) {
	wire.Build(
		ProvideNewComponent01,
		ProvideNewComponent02,

		wire.Struct(new(Component), "*"),
	)
	return &Component{}, nil
}

// 指定使用的参数
func ProvideNewComponent01(config ComponentConfig) component01.Component01er {
	return component01.NewComponent01(config.Component01Params.Field01, config.Component01Params.Field02)
}

// 指定使用的参数
func ProvideNewComponent02(config ComponentConfig) component02.Component02er {
	return component02.NewComponent02(config.Component02Params.Field01, config.Component02Params.Field02)
}
