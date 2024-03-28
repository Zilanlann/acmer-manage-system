package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	// Mongo   Mongo   `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	// Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	// System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	// AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// Local      Local      `mapstructure:"local" json:"local" yaml:"local"`

	// Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
