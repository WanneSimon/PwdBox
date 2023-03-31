package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Debug     bool `yaml:"debug"`     // debug 模式
	Frameless bool `yaml:"frameless"` // 无边框
}

// app 的全局配置对象
var config AppConfig

// app 的全局配置文件
var configPath string = ""

func LoadAppConfig(path string) AppConfig {
	if configPath == "" {
		configPath = path
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("未找到配置文件。" + path)
		log.Println(err)
		config = DefaultAppConfig()
		SaveAppConfig()
		return config
	}

	data, err := ioutil.ReadAll(file)
	if err == nil {
		log.Println("配置文件读取失败 " + path)
		log.Println(err)
		return DefaultAppConfig()
	}

	yaml.Unmarshal(data, config)
	configPath = path
	return config
}

func SaveAppConfig() {
	data, _ := yaml.Marshal(config)
	parent := path.Dir(configPath)
	// parent := filepath.Dir(configPath)

	log.Println("create parent: " + parent)
	os.MkdirAll(parent, 0766)

	log.Println("save config: " + configPath)

	err := ioutil.WriteFile(configPath, data, 0666)
	if err != nil {
		log.Println("配置文件保存失败")
		log.Println(err)
	}
}

func DefaultAppConfig() AppConfig {
	return AppConfig{
		Debug:     false,
		Frameless: false,
	}
}
