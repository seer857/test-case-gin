package utils

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
