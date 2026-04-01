package main

import (
	"log"
	"os"

	slct "github.com/codemityio/auxilium/cmd/select"
	"github.com/codemityio/auxilium/cmd/validate"
	"github.com/codemityio/auxilium/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	application := app.New(
		app.WithValues(
			name,
			`A collection of auxiliary tools for automating and simplifying common project tasks
and supporting project workflows and operations.`,
			version,
			copyright,
			authorName,
			authorEmail,
			buildTime,
		),
	)

	application.Commands = []*cli.Command{
		&slct.App,
		&validate.App,
	}

	if e := application.Run(os.Args); e != nil {
		log.Fatalf("error: %v", e)
	}
}
