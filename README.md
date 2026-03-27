# ![Auxilium](logo.jpg)

![coverage-badge-do-not-edit](https://img.shields.io/badge/Coverage-63%25-yellow.svg?longCache=true&style=flat)

## Table of contents

- [Summary](#summary)
- [Development](#development)
- [Installation](#installation)
- [Usage](#usage)
  - [Manual](#manual)
  - [Subcommands](#subcommands)
  - [Docker](#docker)
- [Packages](#packages)
- [License](#license)

## Summary

A collection of auxiliary tools for automating and simplifying common project tasks and supporting project workflows and
operations.

## Development

To work with the codebase, use `make` command as the primary entry point for all project tools.

Use the arrow keys `↓ ↑ → ←` to navigate the options, and press `/` to toggle search.

## Installation

To install the tool use `make install` (directly from the repository clone) or use
`go install github.com/codemityio/auxilium@latest`.

## Usage

Once you have the tool installed, just use the `auxilium` command to get started.

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
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   codemityio
```

### Subcommands

- [`select`](cmd/select/README.md) - An interactive select command.

### Docker

``` bash
$ docker run codemityio/auxilium
```

## Packages

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
