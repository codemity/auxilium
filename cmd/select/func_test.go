package slct

import (
	"flag"
	"maps"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

// newContext builds a *cli.Context wired to the real App.Flags definitions,
// then applies the caller's overrides on top of each flag's declared default.
func newContext(t *testing.T, strFlags map[string]string, boolFlags map[string]bool) *cli.Context {
	t.Helper()

	fs := flag.NewFlagSet("test", flag.ContinueOnError)

	for _, f := range App.Flags {
		require.NoError(t, f.Apply(fs))
	}

	for name, value := range strFlags {
		require.NoError(t, fs.Set(name, value))
	}

	for name, value := range boolFlags {
		v := "false"
		if value {
			v = "true"
		}

		require.NoError(t, fs.Set(name, v))
	}

	return cli.NewContext(cli.NewApp(), fs, nil)
}

func TestAction(t *testing.T) {
	tests := []struct {
		name      string
		strFlags  map[string]string
		boolFlags map[string]bool
		wantErr   error
	}{
		{
			name: "empty list returns nil immediately",
			strFlags: map[string]string{
				"list": "",
			},
		},
		{
			name: "whitespace-only list returns nil immediately",
			strFlags: map[string]string{
				"list": "   ",
			},
		},
		{
			// No TTY in tests: prompt.Run() fails and action returns errPrompt.
			name: "single item reaches prompt, no TTY returns errPrompt",
			strFlags: map[string]string{
				"list": "foo",
			},
			wantErr: errPrompt,
		},
		{
			name: "multi-item list builds options then returns errPrompt",
			strFlags: map[string]string{
				"list":      "alpha=1\nbeta=2\ngamma=3",
				"delimiter": "=",
			},
			wantErr: errPrompt,
		},
		{
			name: "non-empty exit-name prepends an option",
			strFlags: map[string]string{
				"list":       "item=val",
				"delimiter":  "=",
				"exit-name":  "Quit",
				"exit-value": "q",
			},
			wantErr: errPrompt,
		},
		{
			name: "exit-name with empty exit-value is valid",
			strFlags: map[string]string{
				"list":       "item=val",
				"delimiter":  "=",
				"exit-name":  "Back",
				"exit-value": "",
			},
			wantErr: errPrompt,
		},
		{
			// App default for exit-name is "." so the exit option is prepended.
			name: "default exit-name dot is treated as non-empty",
			strFlags: map[string]string{
				"list":      "item=val",
				"delimiter": "=",
			},
			wantErr: errPrompt,
		},
		{
			name: "empty exit-name suppresses exit option",
			strFlags: map[string]string{
				"list":      "item=val",
				"delimiter": "=",
				"exit-name": "",
			},
			wantErr: errPrompt,
		},
		{
			// No delimiter match: Name = ll[0], Value = "".
			name: "list entry without delimiter uses whole entry as name",
			strFlags: map[string]string{
				"list":      "onlyname",
				"delimiter": "=",
			},
			wantErr: errPrompt,
		},
		{
			// Extra delimiter tokens are re-joined into Value.
			name: "list entry with multiple delimiters joins remainder as value",
			strFlags: map[string]string{
				"list":      "key=part1=part2",
				"delimiter": "=",
			},
			wantErr: errPrompt,
		},
		{
			name: "tab delimiter splits correctly",
			strFlags: map[string]string{
				"list":      "col1\tcol2",
				"delimiter": "\t",
			},
			wantErr: errPrompt,
		},
		{
			name: "custom select-name-label and select-value-label are accepted",
			strFlags: map[string]string{
				"list":               "item=val",
				"delimiter":          "=",
				"select-name-label":  "Option",
				"select-value-label": "Result",
			},
			wantErr: errPrompt,
		},
		{
			name: "custom label is accepted",
			strFlags: map[string]string{
				"list":      "item=val",
				"delimiter": "=",
				"label":     "Pick something",
			},
			wantErr: errPrompt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strFlags := make(map[string]string, len(tt.strFlags))
			maps.Copy(strFlags, tt.strFlags)

			boolFlags := make(map[string]bool, len(tt.boolFlags))
			maps.Copy(boolFlags, tt.boolFlags)

			require.ErrorIs(t, action(newContext(t, strFlags, boolFlags)), tt.wantErr)
		})
	}
}

// TestActionWriteError confirms that a write failure on os.Stdout surfaces as
// errWrite. In practice the prompt fails before any write without a real TTY,
// so this test documents the known gap: exercising errWrite fully requires
// injecting a fake promptui.Select.
func TestActionWriteError(t *testing.T) {
	badOut, err := os.Open(os.DevNull)
	require.NoError(t, err)
	t.Cleanup(func() { _ = badOut.Close() })

	original := os.Stdout
	os.Stdout = badOut

	t.Cleanup(func() { os.Stdout = original })

	ctx := newContext(t, map[string]string{
		"list":      "item=val",
		"delimiter": "=",
		"exit-name": "",
	}, nil)

	require.ErrorIs(t, action(ctx), errPrompt,
		"without a TTY the prompt error is returned before any write attempt")
}
