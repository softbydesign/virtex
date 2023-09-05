package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const DEFAULT_VIRTEX_CONFIG_PATH = "virtex-config.yml"

type VirtexConfig struct {
	Virtex struct {
		Version string `yaml:"version"`

		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`

			LibvirtConnection struct {
				Scheme string `yaml:"scheme"`
				Host   string `yaml:"host"`
				Port   string `yaml:"port"`
			} `yaml:"libvirt-connection"`
		} `yaml:"server"`
	} `yaml:"virtex"`
}

func LoadVirtexConfig(configSourcePath string) (*VirtexConfig, error) {
	var virtexConfig VirtexConfig

	fmt.Printf("Loading Virtex config from %s\n", configSourcePath)

	configFile, err := os.Open(configSourcePath)

	if err != nil {
		fmt.Println("Unable to load Virtex config source")
	}

	defer configFile.Close()

	yamlDecoder := yaml.NewDecoder(configFile)

	if err := yamlDecoder.Decode(&virtexConfig); err != nil {
		return nil, err
	}

	return &virtexConfig, nil
}
