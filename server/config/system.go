package config

// System 基础配置
type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                              // 项目名称
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                              // 项目模式
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                              // 项目使用端口
	Version       string `mapstructure:"version" json:"version" yaml:"version"`                     // 项目版本
	StartTime     string `mapstructure:"start_time" json:"start_time" yaml:"start_time"`            // 项目开始时间
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql   其中sqlite|sqlserver|postgresql 暂未支持 😄
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
}
