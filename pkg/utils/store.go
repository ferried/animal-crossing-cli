package utils

import (
	"io/ioutil"
)

func Store(fileName string) []byte {
	data, err := ioutil.ReadFile("api/" + fileName)
	if nil == err {
		return data
	}
	return nil
}
