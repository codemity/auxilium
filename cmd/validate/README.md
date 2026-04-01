# `validate`

## Table of contents

- [Summary](#summary)
- [Manual](#manual)
- [Usage](#usage)

## Summary

Command line **JSON Schema** validator.

## Manual

``` bash
$ auxilium validate --help
NAME:
   auxilium validate

USAGE:
   auxilium validate [command options]

DESCRIPTION:
   JSON schema validator

OPTIONS:
   --input value         input to be validated
   --input-path input    input file path to be validated (used if input flag is not provided)
   --input-format json   input format to be used (json, `yaml` or `yml`) (default: "json")
   --schema value        schema to validate against
   --schema-path schema  schema file path to validate against (used if schema flag is not provided)
   --schema-format json  schema format to be used (json, `yaml` or `yml`) (default: "json")
   --output-format json  output format to be used (json, `yaml` or `yml`) (default: "json")
   --help, -h            show help
```

## Usage

It supports either **JSON** or **YAML**.

``` shell
auxilium validate \
  --input-path=input.json \
  --schema-path=schema.json
```

``` shell
auxilium validate \
  --input-path=input.yaml \
  --input-format=yaml \
  --schema-path=schema.json
```

``` shell
auxilium validate \
  --input="$(cat input.json)" \
  --schema-path=schema.json
```

``` shell
auxilium validate \
  --input="$(cat input.yaml)" \
  --input-format=yaml \
  --schema-path=schema.json
```
