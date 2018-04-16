package main

import (
	"os"

	"github.com/urfave/cli"
)

var Version string = "0.0.1"

func main() {
    
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "url2mdlink"
	app.Usage = "Convert a url to the markdown link"
	app.Version = Version
	app.Author = "takatori"
	app.Email = "takatori@gmail.com"
	app.Flags = Flags
	app.Action = Action
	return app
}
