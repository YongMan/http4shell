package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	CmdReadOnly = iota
	CmdWrite
	CmdDelete
)

type Config struct {
	Listen string `yaml:"listen"`
	Secret string `yaml:"secret"`
	Cmds   []Cmds `yaml:"cmds"`
}

type Cmds struct {
	Cmd struct {
		Type int    `yaml:"type"`
		Arg  string `yaml:"arg"`
		Desc string `yaml:"desc"`
	}
}

func LoadConfig(file string) (*Config, error) {
	var c Config

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	//unmarshal yaml
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
