package slct

import (
	"github.com/urfave/cli/v2"
)

type input struct {
	Name  string
	Value string
}

var App = cli.Command{ //nolint:gochecknoglobals,exhaustruct
	Name:         "select",
	Aliases:      nil,
	Usage:        "",
	UsageText:    "",
	Description:  "Configurable shell select",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before:       nil,
	After:        nil,
	Action:       action,
	OnUsageError: nil,
	Flags: []cli.Flag{
		&cli.StringFlag{ //nolint:exhaustruct
			Name:     "list",
			Usage:    "a list of options (each option as a new line)",
			Value:    "",
			Required: true,
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "label",
			Usage: "a label to be used in the select",
			Value: "Choose",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "delimiter",
			Usage: "an option delimiter used to split option names from the remaining content to be used as description",
			Value: " ",
		},
		&cli.IntFlag{ //nolint:exhaustruct
			Name:  "size",
			Usage: "a list size",
			Value: selectSize,
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "select-name-label",
			Usage: "a label for the details menu name field label",
			Value: "Name",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "select-value-label",
			Usage: "a label for the details menu value field label",
			Value: "Value",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "exit-name",
			Usage: "a name to be used for the exit option",
			Value: ".",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "exit-value",
			Usage: "a value to be returned for the exit option",
			Value: "",
		},
	},
	Subcommands: nil,
}
