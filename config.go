package jccclient

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ClientConfig
type ClientConfig struct {
	ServAddr string
	Token    string
	Tags     []string
}

// Config - config
type Config struct {

	//------------------------------------------------------------------
	// configuration

	Clients []ClientConfig
}

func checkClientConfig(cfg *ClientConfig) error {
	if cfg.ServAddr == "" {
		return ErrNoServAddr
	}

	if cfg.Token == "" {
		return ErrNoToken
	}

	return nil
}

func checkConfig(cfg *Config) error {
	for _, v := range cfg.Clients {
		err := checkClientConfig(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
