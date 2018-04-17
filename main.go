package main

import (
	"os"

	"github.com/urfave/cli"
)

var Version string = "0.0.2"

func main() {
    
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "mli"
	app.Usage = "Convert a URL to its markdown link"
	app.Version = Version
	app.Author = "takatori"
	app.Email = "takatori@gmail.com"
	app.Flags = Flags
	app.Action = Action
	return app
}
