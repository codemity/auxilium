# ![Auxilium](logo.jpg)

![coverage-badge-do-not-edit](https://img.shields.io/badge/Coverage-64%25-yellow.svg?longCache=true&style=flat)

## Table of contents

- [Summary](#summary)
- [Installation](#installation)
- [Usage](#usage)
  - [Manual](#manual)
  - [Subcommands](#subcommands)
  - [Docker](#docker)
- [Development](#development)
- [Packages](#packages)
- [Dependencies](#dependencies)
  - [Graph](#graph)
  - [Licenses](#licenses)
- [License](#license)

## Summary

A collection of auxiliary tools for automating and simplifying common project tasks and supporting project workflows and
operations.

## Installation

To install the package, run `make install` (directly from the repository clone) or use
`go install github.com/codemityio/auxilium@latest`.

## Usage

Once installed, use the `auxilium` command to get started.

### Manual

``` bash
$ auxilium --help
NAME:
   auxilium - A new cli application

USAGE:
   auxilium [global options] command [command options]

VERSION:
   latest

DESCRIPTION:
   A collection of auxiliary tools for automating and simplifying common project tasks
   and supporting project workflows and operations.

AUTHOR:
   codemityio

COMMANDS:
   select    
   validate  
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   codemityio
```

### Subcommands

- [`select`](cmd/select/README.md) - An interactive select command.
- [`validate`](cmd/validate/README.md) - Command line **JSON Schema** validator.

### Docker

``` bash
$ docker run codemityio/auxilium
```

## Development

To work with the codebase, use `make` command as the primary entry point for all project tools.

Use the arrow keys `↓ ↑ → ←` to navigate the options, and press `/` to toggle search.

## Packages

## Dependencies

### Graph

![](docs/depgraph.svg)

### Licenses

| Package                                 | Licence                                                                             | Type         |
|-----------------------------------------|-------------------------------------------------------------------------------------|--------------|
| github.com/chzyer/readline              | https://github.com/chzyer/readline/blob/2972be24d48e/LICENSE                        | MIT          |
| github.com/cpuguy83/go-md2man/v2/md2man | https://github.com/cpuguy83/go-md2man/blob/v2.0.7/LICENSE.md                        | MIT          |
| github.com/ghodss/yaml                  | https://github.com/ghodss/yaml/blob/v1.0.0/LICENSE                                  | MIT          |
| github.com/manifoldco/promptui          | https://github.com/manifoldco/promptui/blob/v0.9.0/LICENSE.md                       | BSD-3-Clause |
| github.com/russross/blackfriday/v2      | https://github.com/russross/blackfriday/blob/v2.1.0/LICENSE.txt                     | BSD-2-Clause |
| github.com/urfave/cli/v2                | https://github.com/urfave/cli/blob/v2.27.7/LICENSE                                  | MIT          |
| github.com/xeipuuv/gojsonpointer        | https://github.com/xeipuuv/gojsonpointer/blob/4e3ac2762d5f/LICENSE-APACHE-2.0.txt   | Apache-2.0   |
| github.com/xeipuuv/gojsonreference      | https://github.com/xeipuuv/gojsonreference/blob/bd5ef7bd5415/LICENSE-APACHE-2.0.txt | Apache-2.0   |
| github.com/xeipuuv/gojsonschema         | https://github.com/xeipuuv/gojsonschema/blob/v1.2.0/LICENSE-APACHE-2.0.txt          | Apache-2.0   |
| github.com/xrash/smetrics               | https://github.com/xrash/smetrics/blob/686a1a2994c1/LICENSE                         | MIT          |
| gopkg.in/yaml.v2                        | https://github.com/go-yaml/yaml/blob/v2.4.0/LICENSE                                 | Apache-2.0   |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
