package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	defer fmt.Println("first defer func")
	defer file.Close()
	defer fmt.Println("last defer func")
	if err != nil {
		panic(err)
		// return "", err
	}
	bys, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bys), nil
}
