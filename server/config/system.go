package config

// System 基础配置
type System struct {
	Name      string `mapstructure:"name" json:"name" yaml:"name"`
	Mode      string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Port      int    `mapstructure:"port" json:"port" yaml:"port"`
	Version   string `mapstructure:"version" json:"version" yaml:"version"`
	StartTime string `mapstructure:"start_time" json:"start_time" yaml:"start_time"`
	DbType    string `mapstructure:"db-type" json:"dbType" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
