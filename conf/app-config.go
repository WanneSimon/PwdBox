package conf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Debug     bool `yaml:"debug"`     // debug 模式
	Frameless bool `yaml:"frameless"` // 无边框

	Emojis []string `yaml:"emojis"` // 图片路径
}

// app 的全局配置对象
var Config AppConfig

// app 的全局配置文件
var ConfigPath string = ""

// //go:embed saya-sample.yml
// var sayaConfigString string

func LoadAppConfig(path string) AppConfig {
	log.Println("读取配置文件: " + path)

	file, err := os.Open(path)
	if err != nil {
		log.Println("未找到配置文件")
		log.Println(err)
		ConfigPath = path
		Config = DefaultAppConfig()
		Config.SaveAppConfig()
		return Config
	}

	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Println("配置文件读取失败")
		log.Println(err2)
		Config = DefaultAppConfig()
		Config.SaveAppConfig()
		return Config
	}

	err3 := yaml.Unmarshal(data, &Config)
	if err3 != nil {
		log.Fatalln(err3)
	}

	ConfigPath = path
	return Config
}

func (ac *AppConfig) SaveAppConfig() bool {
	data, _ := yaml.Marshal(ac)
	// parent := path.Dir(ConfigPath)
	parent := filepath.Dir(ConfigPath)

	log.Println("config path: " + ConfigPath)
	log.Println("create parent: " + parent)
	log.Println(ac)
	os.MkdirAll(parent, 0766)

	err := ioutil.WriteFile(ConfigPath, data, 0666)
	if err != nil {
		log.Println("配置文件保存失败")
		log.Println(err)
		return false
	}

	log.Println("配置文件保存成功")
	return true
}

func DefaultAppConfig() AppConfig {
	return AppConfig{
		Debug:     false,
		Frameless: true,
		Emojis:    []string{},
	}
}

// func (ac *AppConfig) Save() bool {
// 	return SaveAppConfig()
// }
