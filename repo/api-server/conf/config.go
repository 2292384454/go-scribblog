package conf

import (
	"go-scribblog/repo/log"
	"gopkg.in/yaml.v3"
	"os"
)

type Server struct {
	Name    string `yaml:"web_name"`
	WebAddr string `yaml:"web_addr"`
}

type Config struct {
	Server Server      `yaml:"server"`
	Log    log.Options `yaml:"log"`
}

func LoadConfigFromFile(configPath string) (*Config, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(content, config)
	return config, err
}
