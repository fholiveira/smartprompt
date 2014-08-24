package git

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
)

type FileReader struct{}

func (reader FileReader) ReadFirstLine(path string) (string, error) {
	file, err := os.Open(path)
	if nil != err {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("Empty file")
}

func (reader FileReader) ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if nil != err {
		return "", err
	}

	return string(content), nil
}

func (reader FileReader) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if nil != err {
		return false, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
