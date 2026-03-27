package slct

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	list := ctx.String("list")
	label := ctx.String("label")
	delimiter := ctx.String("delimiter")
	selectNameLabel := ctx.String("select-name-label")
	selectValueLabel := ctx.String("select-value-label")
	exitName := ctx.String("exit-name")
	exitValue := ctx.String("exit-value")

	if strings.TrimSpace(list) == "" {
		return nil
	}

	lst := strings.Split(list, "\n")

	var options []input

	if exitName != "" {
		options = append(options, input{
			Name:  exitName,
			Value: exitValue,
		})
	}

	for i := range lst {
		ll := strings.Split(lst[i], delimiter)

		options = append(options, input{
			Name:  ll[0],
			Value: strings.Join(ll[1:], delimiter),
		})
	}

	templates := &promptui.SelectTemplates{ //nolint:exhaustruct
		Label:    "{{ . | yellow }}{{ \":\" | yellow }}",
		Active:   "> {{ .Name | cyan | bold }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "{{ \"Selected:\" | bold }} {{ .Name }}",
		Details: fmt.Sprintf(`
{{ "Select:" | yellow }}
{{ "%s:" | faint }} {{ .Name }}
{{ if .Value }}{{ "%s:" | faint }} {{ .Value }}{{ end }}`, selectNameLabel, selectValueLabel),
	}

	searcher := func(input string, index int) bool {
		option := options[index]
		name := strings.ReplaceAll(strings.ToLower(option.Name), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{ //nolint:exhaustruct
		Label:     label,
		Items:     options,
		Templates: templates,
		Size:      selectSize,
		Searcher:  searcher,
		Stdout:    os.Stderr,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("%w: %w", errPrompt, err)
	}

	if index == 0 {
		if _, e := os.Stdout.WriteString(options[index].Value); e != nil {
			return fmt.Errorf("%w: %w", errWrite, e)
		}

		return nil
	}

	var output string

	if options[index].Value == "" {
		output = options[index].Name
	}

	if options[index].Value != "" {
		output = options[index].Name + delimiter + options[index].Value
	}

	if _, e := os.Stdout.WriteString(output); e != nil {
		return fmt.Errorf("%w: %w", errWrite, e)
	}

	return nil
}
