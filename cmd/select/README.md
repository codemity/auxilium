# `select`

## Table of contents

- [Summary](#summary)
- [Manual](#manual)
- [Usage](#usage)
  - [Example](#example)

## Summary

An interactive select command.

## Manual

``` bash
$ auxilium select --help
NAME:
   auxilium select

USAGE:
   auxilium select [command options]

DESCRIPTION:
   Configurable shell select

OPTIONS:
   --list value                a list of options (each option as a new line)
   --label value               a label to be used in the select (default: "Choose")
   --delimiter value           an option delimiter used to split option names from the remaining content to be used as description (default: " ")
   --size value                a list size (default: 15)
   --select-name-label value   a label for the details menu name field label (default: "Name")
   --select-value-label value  a label for the details menu value field label (default: "Value")
   --exit-name value           a name to be used for the exit option (default: ".")
   --exit-value value          a value to be returned for the exit option
   --help, -h                  show help
```

## Usage

Coming soon…

### Example

``` shell
list=$(cat <<-EOF
one one-value-one one-value-two
two two-value-one two-value-two
EOF
)
auxilium select --label="Choose one of the following" --list="${list}"
```
