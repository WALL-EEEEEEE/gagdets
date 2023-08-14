package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
}

func parseYML(fname string, cfg *Config) (err error) {
	f, err := os.Open(fname)
	if err != nil {
		return fmt.Errorf("Failed to open file %s", fname)
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return fmt.Errorf("file %s is not an valid yaml file", fname)
	}
	return nil
}

func parseEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error(err)
		return
	}
}

func ParseConfig(f string) (cfg Config, err error) {
	err = parseYML(f, &cfg)
	if err != nil {
		return cfg, err
	}
	parseEnv(&cfg)
	return cfg, nil
}
