package main

import (
	"github.com/s-shin/gobbs/db"
	"log"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
}
