package config

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type ServerConfig struct {
	Port string `yaml:"APP_PORT"`
	Mode string `yaml:"MODE"`
}

type LoggerConfig struct {
	Level        string `yaml:"LEVEL"`
	ErrorOutPath string `yaml:"ERR_OUT"`
	AllOutPath   string `yaml:"ALL_OUT"`
}

func (c LoggerConfig) GetLevel() logrus.Level {
	switch c.Level {
	case "PANIC":
		return logrus.PanicLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "WARNING":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	}
	return logrus.TraceLevel
}

type ServiceConfig struct {
	ServerConfig ServerConfig `yaml:"server"`
	LoggerConfig LoggerConfig `yaml:"logger"`
}

const DefaultPath = "conf.yaml"

func (c ServiceConfig) FromFile() ServiceConfig {
	var path string
	flag.StringVar(&path, "c", DefaultPath, "path to conf")
	flag.Parse()
	filename, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Sprintf("can't get file path %s", path))
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("can't open file %s", filename))
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	return c
}
