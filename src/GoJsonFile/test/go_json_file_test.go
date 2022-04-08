package test

import (
	"log"
	"path"
	"testing"
)

func Test(t *testing.T) {
	str := GoJsonFile.pwd()
	log.Printf("%v\n", str)

	str = path.Join(str, "./config/database.json")
	log.Printf("%v\n", str)
}
