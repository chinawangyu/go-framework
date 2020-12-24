package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"go-framework/http/pkg/mysql"
	"io/ioutil"
	"log"
	"os"
)

var Config Configure

type Configure struct {
	Env  string
	Mode string
	Port int

	MysqlMaster mysql.DbConfig
	MysqlSlave  mysql.DbConfig
}

func Init() error {
	projectPath, _ := os.Getwd()
	configpath := flag.String("f", projectPath+"/http/config/config.toml", "config file")
	flag.Parse()
	err := initConfig(*configpath)

	log.Println(GetMode())

	if err != nil {
		return err
	}
	return nil
}

func initConfig(configpath string) error {
	configBytes, err := ioutil.ReadFile(configpath)
	if err != nil {
		return err
	}
	if _, err := toml.Decode(string(configBytes), &Config); err != nil {
		return err
	}
	log.Println(configpath)

	return nil
}

func GetMode() string {
	return Config.Mode
}

func GetPort() string {
	return fmt.Sprintf(":%d", Config.Port)
}
