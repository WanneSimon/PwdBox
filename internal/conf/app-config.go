package conf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// 只能暴露给前端结构体的函数，成员变量不能访问，只能通过函数获取
// 推荐： 函数 和 数据使用不同的结构体
type AppConfig struct {
	Title     string `yaml:"title" json:"Title"`         // 程序名称
	Debug     bool   `yaml:"debug" json:"Debug"`         // debug 模式
	Frameless bool   `yaml:"frameless" json:"Frameless"` // 无边框

	Pwdbox string `yaml:"pwdbox" json:"pwdbox"` // 保险箱路径
}

// app 的全局配置对象
var Config AppConfig

// app 的全局配置文件
var ConfigPath string = ""

// //go:embed saya-sample.yml
// var sayaConfigString string

// 设置全局配置对象
func SetConfig(ac AppConfig) {
	Config = ac
}

func DefaultAppConfig() AppConfig {
	return AppConfig{
		Title:     "saya-app",
		Debug:     false,
		Frameless: true,
		Pwdbox:    "config/pwdbox.db3",
	}
}

// 暴露给外部的操作对象
type ConfigOps struct {
}

// 初始化空对象
func NewConfigOps() *ConfigOps {
	return &ConfigOps{}
}

// 初始化并读取配置文件
func NewConfigOpsAndLoad(path string) *ConfigOps {
	co := ConfigOps{}
	co.LoadAppConfig(path)
	return &co
}

// 不要在前端调用这个函数，前端的数据修改后并不会同步到后端
func (co *ConfigOps) Save() bool {
	// fmt.Println(ac)
	data, _ := yaml.Marshal(Config)
	// parent := path.Dir(ConfigPath)
	parent := filepath.Dir(ConfigPath)

	log.Println("config path: " + ConfigPath)
	log.Println("create parent: " + parent)
	log.Println(Config)
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

// 这个是给前端调用的，修改全局配置并保存
func (co *ConfigOps) SaveAppConfig(newAc AppConfig) bool {
	// data, _ := yaml.Marshal(ac)
	data, _ := yaml.Marshal(newAc)
	// parent := path.Dir(ConfigPath)
	parent := filepath.Dir(ConfigPath)

	log.Println("config path: " + ConfigPath)
	log.Println("create parent: " + parent)
	log.Println(newAc)
	os.MkdirAll(parent, 0766)

	err := ioutil.WriteFile(ConfigPath, data, 0666)
	if err != nil {
		log.Println("配置文件保存失败")
		log.Println(err)
		return false
	}

	Config = newAc
	log.Println("配置文件保存成功")
	return true
}

// 刷新
func (co *ConfigOps) RefreshConfig() AppConfig {
	return co.LoadAppConfig(ConfigPath)
}

// 加载
func (co *ConfigOps) LoadAppConfig(path string) AppConfig {
	log.Println("读取配置文件: " + path)

	file, err := os.Open(path)
	if err != nil {
		log.Println("未找到配置文件")
		log.Println(err)
		ConfigPath = path
		Config = DefaultAppConfig()
		co.SaveAppConfig(Config)
		return Config
	}

	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Println("配置文件读取失败")
		log.Println(err2)
		Config = DefaultAppConfig()
		co.SaveAppConfig(Config)
		return Config
	}

	err3 := yaml.Unmarshal(data, &Config)
	if err3 != nil {
		log.Fatalln(err3)
	}

	if Config.Title == "" {
		Config.Title = "Saya"
	}

	ConfigPath = path
	return Config
}

func (co *ConfigOps) Get() *AppConfig {
	return &Config
}

// func (ac *AppConfig) Save() bool {
// 	return SaveAppConfig()
// }
