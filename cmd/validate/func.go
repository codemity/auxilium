package validate

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli/v2"
	"github.com/xeipuuv/gojsonschema"
)

func action(ctx *cli.Context) error {
	inputContent, err := loadContent(ctx.String("input"), ctx.String("input-path"))
	if err != nil {
		return err
	}

	schemaContent, err := loadContent(ctx.String("schema"), ctx.String("schema-path"))
	if err != nil {
		return err
	}

	subject, err := unmarshal(inputContent, ctx.String("input-format"))
	if err != nil {
		return err
	}

	spec, err := unmarshal(schemaContent, ctx.String("schema-format"))
	if err != nil {
		return err
	}

	errors, err := validate(spec, subject)
	if err != nil {
		return err
	}

	if len(errors) > 0 {
		if e := writeErrors(errors, ctx.String("output-format")); e != nil {
			return e
		}

		os.Exit(1)
	}

	return nil
}

func loadContent(inline, path string) ([]byte, error) {
	if inline != "" {
		return []byte(inline), nil
	}

	content, err := os.ReadFile(path) //#nosec G304
	if err != nil {
		return nil, fmt.Errorf("%w: incorrect file path `%s`: %w", errRead, path, err)
	}

	return content, nil
}

func unmarshal(content []byte, frmt string) (any, error) {
	var out any

	var err error

	switch format(frmt) {
	case jsonFmt:
		err = json.Unmarshal(content, &out)
	case yamlFmt, ymlFmt:
		err = yaml.Unmarshal(content, &out)
	default:
		return nil, fmt.Errorf("%w: with `%s` format", errFormat, frmt)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errMarshal, err)
	}

	return out, nil
}

func validate(spec, subject any) ([]validationError, error) {
	res, err := gojsonschema.Validate(
		gojsonschema.NewGoLoader(spec),
		gojsonschema.NewGoLoader(subject),
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errValidate, err)
	}

	var errors []validationError

	for _, d := range res.Errors() {
		errors = append(errors, validationError{
			Field:   d.Field(),
			Message: d.Description(),
		})
	}

	return errors, nil
}

func writeErrors(errors []validationError, frmt string) error {
	var (
		out []byte
		err error
	)

	switch format(frmt) {
	case jsonFmt:
		out, err = json.Marshal(errors)
	case yamlFmt, ymlFmt:
		out, err = yaml.Marshal(errors)
	default:
		return fmt.Errorf("%w: with `%s` output format", errFormat, frmt)
	}

	if err != nil {
		return fmt.Errorf("%w: %w", errMarshal, err)
	}

	if _, e := os.Stdout.WriteString(string(out)); e != nil {
		return fmt.Errorf("%w: %w", errWrite, e)
	}

	return nil
}
