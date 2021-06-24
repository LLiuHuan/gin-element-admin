package config

type RateLimit struct {
	IpVerify   bool   `mapstructure:"ip-verify" json:"ip-verify" yaml:"ip-verify"`          // 是否打开ip限流
	IpLimitCon int64  `mapstructure:"ip-limit-con" json:"ip-limit-con" yaml:"ip-limit-con"` // 每秒访问ip超过多少次
	IpListKey  string `mapstructure:"ip-list-key" json:"ip-list-key" yaml:"ip-list-key"`    // ip列表的key

	Cap     int64 `mapstructure:"cap" json:"cap" yaml:"cap"`             // 初始化数量
	Quantum int64 `mapstructure:"quantum" json:"quantum" yaml:"quantum"` // 每秒增加数量
}
