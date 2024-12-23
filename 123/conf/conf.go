package conf

import (
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env    string
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	
	Redis    Redis    `yaml:"redis"`
	Logger   Logger   `yaml:"logger"`
	Cors     Cors     `yaml:"cors"`
}

type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type Logger struct {
	FilePath string `yaml:"filePath"`
	Encoding string `yaml:"encoding"`
	Level    string `yaml:"level"`
	Logger   string `yaml:"logger"`
}




type Mysql struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DbName          string        `yaml:"dbName"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}



type Redis struct {
	Host               string        `yaml:"host"`
	Port               int           `yaml:"port"`
	DB                 int           `yaml:"db"`
	DialTimeout        time.Duration `json:"dialTimeout"`
	ReadTimeout        time.Duration `json:"readTimeout"`
	WriteTimeout       time.Duration `json:"writeTimeout"`
	IdleCheckFrequency time.Duration `json:"idleCheckFrequency"`
	PoolSize           int           `json:"poolSize"`
	PoolTimeout        time.Duration `json:"poolTimeout"`
}


type Cors struct {
	AllowOrigins string `yaml:"allowOrigins"`
}

func GetConfig() *Config {
	initConf()
	return conf
}

func initConf() {
	once.Do(func() {
		dir := "conf"
		confFileRelPath := filepath.Join(dir, filepath.Join(GetEnv(), "config.yaml"))
		confContent, err := os.ReadFile(confFileRelPath)
		if err != nil {
			panic(err)
		}
		conf = new(Config)
		err = yaml.Unmarshal(confContent, conf)
		if err != nil {
			panic(err)
		}
		if err := validator.Validate(conf); err != nil {
			panic(err)
		}
		conf.Env = GetEnv()
		pretty.Printf("%+v\n", conf)
	})
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "local"
	}
	return e
}

