package common

import (
	"gopkg.in/ini.v1"
	"log"
)

type GlobalConfig struct { //配置文件要通过tag来指定配置文件中的名称
	RediszzHostname string `ini:"rediszz_hostname"`
	RediszzPort     string `ini:"rediszz_port"`
	RediszzPass     string `ini:"rediszz_pass"`
	RediszzTimeout  int    `ini:"rediszz_timeout"`

	LogPath string `ini:"logpath"`
	LogFile string `ini:"logfile"`
}

//加载配置文件
func LoadConf() (GlobalConfig, error) {
	var config GlobalConfig
	conf, err := ini.Load(ConfigPath) //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
