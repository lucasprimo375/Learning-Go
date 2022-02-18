package app

import "github.com/urfave/cli"

// Generate returns the command line application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Searches for IPs and servers names on the web"

	return app
}
