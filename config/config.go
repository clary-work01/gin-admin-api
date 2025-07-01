package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// 總配置文件
type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
}
// 系統配置
type system struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env string `yaml:"env"`
} 
// 日誌配置
type logger struct {
	Level string `yaml:"level"`
	Prefix string `yaml:"prefix"`
	Director string `yaml:"director"`
	ShowLine bool `yaml:"show_line"`
	LogInConsole bool `yaml:"log_in_console"`
}
// mysql配置
type mysql struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Db string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
	Charset string `yaml:"charset"`
	MaxIdle int `yaml:"max_idle"`
	MaxOpen int `yaml:"max_open"`
}
// redis配置
type redis struct {
	Address string `yaml:"address"`
	Password string `yaml:"password"`
	Db int `yaml:"db"`
}

var Config *config

func init (){
	yamlFile,err:=os.ReadFile("./config.yaml")
	if err != nil {
		return 
	}
	yaml.Unmarshal(yamlFile, &Config)
}