package setup

import (
	"log"
	"os"

	"github.com/morgann-erik/sb/core"
	"github.com/urfave/cli/v2"
)

func SetupCommand(ctx *cli.Context) error {
	log.Println("Setup command")
	if err := valiateDataDir(); err != nil {
		return err
	}

	return nil
}

func valiateDataDir() error {
	path, err := core.DataDirPath()
	if err != nil {
		return err
	}

	log.Println(path)

	// if exists early out
	_, err = os.Stat(path)
	if os.IsExist(err) {
		return nil
	}

	// create dir
	os.Mkdir(path, os.ModePerm)

	return nil
}
