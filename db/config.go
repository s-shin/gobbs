package db

import (
	"github.com/s-shin/gobbs/config"
	"github.com/s-shin/gobbs/util"
	"log"
	"path"
)

func GetDSN(name string) string {
	c := config.Config
	var p string
	if name == "master" {
		p = c.DB.Master
	} else if name == "slave" {
		p = c.DB.Slave
	} else {
		log.Fatalf("Unkown name: %s", name)
	}
	return path.Join(util.CWD(), p)
}
