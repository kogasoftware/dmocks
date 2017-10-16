package configuration

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Configuration is struct of config file.
type Configuration struct {
	RootPath   string      `yaml:"root_path"`
	WriteFiles []WriteFile `yaml:"write_files"`
}

// WriteFile is struct of write_files column in config file.
type WriteFile struct {
	Path     string   `yaml:"path"`
	Interval int      `yaml:"interval"`
	Data     []string `yaml:"data"`
}

// NewConfiguration returns Configuration
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// LoadConfigFile laods config file.
func (config *Configuration) LoadConfigFile(configPath string) (string, error) {
	buffer, err := ioutil.ReadFile(configPath)

	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(buffer, config)

	if err != nil {
		return "", err
	}

	d, err := yaml.Marshal(config)

	return string(d), nil
}
