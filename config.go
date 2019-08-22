package jccclient

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// ClientConfig -
type ClientConfig struct {
	ServAddr string
	Token    string
	Tags     []string
}

// Config - config
type Config struct {
	//------------------------------------------------------------------
	// clients

	Clients []ClientConfig

	//------------------------------------------------------------------
	// configuration

	// MaxMultiFailTimes - If the client fail {{MaxMultiFailTimes}} times in succession, it need to ignore some task
	MaxMultiFailTimes int
	// EffectiveFailTime - If the interval between 2 failures is less than {{EffectiveFailTime}} seconds, it should be counted as a multiple failure.
	EffectiveFailTime int
	// IgnoreTaskTime - The client will ignore the task for {{IgnoreTaskTime}} seconds.
	IgnoreTaskTime int
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
	mapServAddr := make(map[string]bool)

	for _, v := range cfg.Clients {
		_, isok := mapServAddr[v.ServAddr]
		if isok {
			return ErrDuplicateServAddr
		}

		err := checkClientConfig(&v)
		if err != nil {
			return err
		}

		mapServAddr[v.ServAddr] = true
	}

	if cfg.MaxMultiFailTimes <= 0 {
		cfg.MaxMultiFailTimes = MaxMultiFailTimes
	}

	if cfg.EffectiveFailTime <= 0 {
		cfg.EffectiveFailTime = EffectiveFailTime
	}

	if cfg.IgnoreTaskTime <= 0 {
		cfg.IgnoreTaskTime = IgnoreTaskTime
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
