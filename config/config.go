package config

import (
	"github.com/s-shin/gobbs/bindata"
	// "github.com/s-shin/gobbs/util"
	"gopkg.in/yaml.v2"
	// "io/ioutil"
	"log"
	"path"
)

type config_ struct {
	DB struct {
		Master string
		Slave  string
	}
}

var Config *config_

func init() {
	data, err := bindata.Asset(path.Join("config", MODE+".yml"))
	if err != nil {
		log.Fatal(err)
	}
	var c config_
	yaml.Unmarshal(data, &c)
	Config = &c
}
