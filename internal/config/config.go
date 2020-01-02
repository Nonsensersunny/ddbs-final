package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type HttpConf struct {
	Port        int      `yaml:"port"`
	Host        string   `yaml:"host"`
	AllowOrigin []string `yaml:"allow-origin"`
}

type MySQLConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MySQLConfig *MySQLConf

type MySQLs struct {
	MySQLConfig MySQLConfig `yaml:"mysql-config"`
	Area        string      `yaml:"area"`
}

type MySQLsConfig []*MySQLs

type RedisConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   int    `yaml:"dbname"`
}

type ServerConf struct {
	MySQLDBs MySQLsConfig `yaml:"mysqls"`
	RedisDB  *RedisConf   `yaml:"redis"`
	Http     *HttpConf    `yaml:"http"`
}

func GetServerConf() *ServerConf {
	var serverConf *ServerConf
	ymlFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Loading config file found error: %v", err)
	}
	if err := yaml.Unmarshal(ymlFile, &serverConf); err != nil {
		log.Fatalf("Unmarshal config file found error: %v", err)
	}
	return serverConf
}

func GetMySQLConfigByArea(area string) (*MySQLConf, error) {
	serverConf := GetServerConf()
	for _, db := range serverConf.MySQLDBs {
		if db.Area == area {
			return db.MySQLConfig, nil
		}
	}
	return nil, errors.New("unknown area")
}

func GetRedisConfig() *RedisConf {
	return GetServerConf().RedisDB
}