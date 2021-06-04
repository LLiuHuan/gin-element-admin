package config

type Captcha struct {
	NoiseCount int `mapstructure:"noise-count" json:"noiseCount" yaml:"noise-count"` // 验证码长度
	ImgWidth   int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`       // 验证码宽度
	ImgHeight  int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`    // 验证码高度
}
