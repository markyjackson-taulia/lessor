package main

import (
	"math/rand"
	"time"

	"github.com/kolide/kit/version"
	"github.com/urfave/cli"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	app := cli.NewApp()
	app.Name = "lessor"
	app.Usage = "Managing and deploying single-tenant apps"
	app.Version = version.Version().Version
	cli.VersionPrinter = func(c *cli.Context) {
		version.PrintFull()
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:        "adopt",
			Usage:       "Adopt an existing cluster",
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "create",
			Usage:       "Create resources",
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "get",
			Usage:       "Get and list resources",
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "put",
			Usage:       "Create or update resources",
			Subcommands: []cli.Command{},
		},
		cli.Command{
			Name:        "delete",
			Usage:       "Delete resources",
			Subcommands: []cli.Command{},
		},
	}

	app.RunAndExitOnError()
}