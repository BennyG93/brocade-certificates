package brocade

import (
	"io/ioutil"
	"log"
)

func Readfile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	cert := string(data)
	return cert
}