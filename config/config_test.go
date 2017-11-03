package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	file := "/home/dev/gopath/src/github.com/YongMan/http4shell/config.yml"

	c, err := LoadConfig(file)
	fmt.Println(c, err)
}
