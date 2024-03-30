package setup

import (
	"log"

	"github.com/urfave/cli/v2"
)

func SetupCommand(ctx *cli.Context) error {
	log.Println("Setup command")
    startWizard()

	return nil
}
