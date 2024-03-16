package core

import (
	"os"
	"path"
)

const (
	DataDir string = "data"
)

func BasePath() (string, error) {
	p, err := os.Executable() 
    if err != nil {
		return "", err
	}

	return path.Dir(p), nil
}

func DataDirPath() (string, error) {
	base, err := BasePath()
	path := path.Join(base, DataDir)
	if err != nil {
		return "", err
	}

	return path, err
}