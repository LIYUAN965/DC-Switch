package config

import (
	"dcswitch/internal/static"
	"dcswitch/pkg/encrypt"
	"dcswitch/pkg/mysql"
	"embed"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
)

// GlobalConf 全局配置单例
var GlobalConf Conf

var SwaggerFS = static.SwaggerFS

//go:embed config.yaml
var configFile embed.FS

type Conf struct {
	DataBase struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		DBName   string `yaml:"DBName"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Database"`

	SecKey string `yaml:"SecKey"`
}

func (conf *Conf) InitConfig() *Conf {

	switch runtime.GOOS {
	case "windows":
		_ = os.Setenv("ENV", "DEV")
	case "darwin":
		_ = os.Setenv("ENV", "DEV")
	default:
	}

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		GlobalConf.initDevConfig("internal/config/config.yaml")
		log.Infof("START SERVERS WITH %v CONFIG\n", env)
	default:
		GlobalConf.initProdConfig()
		log.Infof("START SERVERS WITH DEFAULT CONFIG\n")
	}
	return conf
}

func (conf *Conf) initDevConfig(configPath string) *Conf {
	yamlFile, err := configFile.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

func (conf *Conf) initProdConfig() {
	panic("not implemented")
}

func InitDB() {
	encryptedPWD := GlobalConf.DataBase.Password
	secKey := GlobalConf.SecKey
	pwd := encrypt.Decrypt(encryptedPWD, secKey)

	o := mysql.DBConnOptions{
		User:     GlobalConf.DataBase.Username,
		Pwd:      pwd,
		Host:     GlobalConf.DataBase.Host,
		Port:     GlobalConf.DataBase.Port,
		Database: GlobalConf.DataBase.DBName,
	}
	mysql.DB.InitConn(o)
}
