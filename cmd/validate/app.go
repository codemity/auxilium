package validate

import (
	_ "embed"

	"github.com/urfave/cli/v2"
)

// App main application.
var App = cli.Command{ //nolint:gochecknoglobals,exhaustruct
	Name:         "validate",
	Aliases:      nil,
	Usage:        "",
	UsageText:    "",
	Description:  "JSON schema validator",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before:       nil,
	After:        nil,
	Action:       action,
	OnUsageError: nil,
	Flags: []cli.Flag{
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "input",
			Usage: "input to be validated",
			Value: "",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "input-path",
			Usage: "input file path to be validated (used if `input` flag is not provided)",
			Value: "",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "input-format",
			Usage: "input format to be used (`json`, `yaml` or `yml`)",
			Value: "json",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "schema",
			Usage: "schema to validate against",
			Value: "",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "schema-path",
			Usage: "schema file path to validate against (used if `schema` flag is not provided)",
			Value: "",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "schema-format",
			Usage: "schema format to be used (`json`, `yaml` or `yml`)",
			Value: "json",
		},
		&cli.StringFlag{ //nolint:exhaustruct
			Name:  "output-format",
			Usage: "output format to be used (`json`, `yaml` or `yml`)",
			Value: "json",
		},
	},
	Subcommands: []*cli.Command{},
}
