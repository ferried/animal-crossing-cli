package utils

import (
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
)

func Store(fileName string) []byte {
	dir, _ := homedir.Dir()
	data, err := ioutil.ReadFile(dir + "/.acc/" + fileName)
	if nil == err {
		return data
	}
	return nil
}
