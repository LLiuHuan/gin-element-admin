package config

type JWT struct {
	SigningKey         string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`                             // jwt签名
	AccessExpiresTime  int64  `mapstructure:"access-expires-time" json:"access-expires-time" yaml:"access-expires-time"`    // access_token 过期时间
	RefreshExpiresTime int64  `mapstructure:"refresh-expires-time" json:"refresh-expires-time" yaml:"refresh-expires-time"` // access_token 过期时间
	//BufferTime         int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`                             // 缓冲时间
}
