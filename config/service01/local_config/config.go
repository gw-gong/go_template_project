package common_config

type Config struct {
	Component01 Component01Config `yaml:"component01" mapstructure:"component01"`
	Component02 Component02Config `yaml:"component02" mapstructure:"component02"`
}

type Component01Config struct {
	Field01 string `yaml:"field01" mapstructure:"field01"`
	Field02 string `yaml:"field02" mapstructure:"field02"`
}

type Component02Config struct {
	Field01 string `yaml:"field01" mapstructure:"field01"`
	Field02 string `yaml:"field02" mapstructure:"field02"`
}
