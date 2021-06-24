package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql     Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	RateLimit RateLimit `mapstructure:"rate-limit" json:"rate-limit" yaml:"rate-limit"`
}
