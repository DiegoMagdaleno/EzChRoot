package lib

import (
	"fmt"
	"os"
)

func fileExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if file.Mode().IsRegular() != true {
		err = fmt.Errorf("libezchroot: Path %s is not a file, when a file was expected", path)
		return true, err
	}
	if err != nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func dirExists(path string) (bool, error) {
	file, err := os.Stat(path)
	if file.Mode().IsDir() != true {
		err = fmt.Errorf("libezchroot: Path %s is not a directory when directory was expected", path)
		return true, err
	}
	if err != nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
