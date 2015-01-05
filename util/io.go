package util

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) string {
	s, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
