package library

import (
	"gopkg.in/ini.v1"
	"log"
)

type config struct {
	env string
	cfg *ini.File
}

var conf *config

func RegisterConfig(projectPath string) {
	file := projectPath + "/conf/index.ini"
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatal("Fail to read file: " + err.Error())
	}
	conf = &config{
		env: "dev",
		cfg: cfg,
	}
}

func NewConfig() *config {
	if conf == nil {
		log.Fatal("config is not set!")
	}
	return conf
}

const (
	ENV_DEV  = "dev"
	ENV_TEST = "test"
	ENV_PROD = "prod"
)

func (c *config) SetEnv(key string) {
	c.env = key
}

func (c *config) GetString(key string) string {
	v := c.cfg.Section("").Key(key).String()
	if v != "" {
		return v
	}
	return c.cfg.Section(c.env).Key(key).String()
}

func (c *config) GetInt(key string) int {
	v, err := c.cfg.Section("").Key(key).Int()
	if err == nil {
		return v
	}
	v, err = c.cfg.Section(c.env).Key(key).Int()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return v
}

func (c *config) GetStrings(key string) []string {
	arr := c.cfg.Section("").Key(key).Strings(",")
	return arr
}
