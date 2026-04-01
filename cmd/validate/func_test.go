package validate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdata = "testdata"

func TestLoadContent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		inline  string
		path    string
		want    []byte
		wantErr error
	}{
		{
			name:   "inline content takes precedence",
			inline: `{"name": "input"}`,
			path:   "",
			want:   []byte(`{"name": "input"}`),
		},
		{
			name: "reads json from file when inline is empty",
			path: filepath.Join(testdata, "input.json"),
		},
		{
			name: "reads yaml from file when inline is empty",
			path: filepath.Join(testdata, "input.yaml"),
		},
		{
			name:    "returns error for invalid path",
			path:    "nonexistent/file.json",
			wantErr: errRead,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := loadContent(tt.inline, tt.path)

			require.ErrorIs(t, err, tt.wantErr)

			if tt.want != nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	jsonContent, err := os.ReadFile(filepath.Join(testdata, "input.json"))
	require.NoError(t, err)

	yamlContent, err := os.ReadFile(filepath.Join(testdata, "input.yaml"))
	require.NoError(t, err)

	tests := []struct {
		name    string
		content []byte
		format  string
		wantErr error
	}{
		{
			name:    "valid json",
			content: jsonContent,
			format:  "json",
		},
		{
			name:    "valid yaml",
			content: yamlContent,
			format:  "yaml",
		},
		{
			name:    "valid yml",
			content: yamlContent,
			format:  "yml",
		},
		{
			name:    "invalid json returns marshal error",
			content: []byte(`{invalid}`),
			format:  "json",
			wantErr: errMarshal,
		},
		{
			name:    "unsupported format returns format error",
			content: []byte(`anything`),
			format:  "toml",
			wantErr: errFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := unmarshal(tt.content, tt.format)

			require.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				assert.NotNil(t, got)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	schemaContent, err := os.ReadFile(filepath.Join(testdata, "schema.json"))
	require.NoError(t, err)

	spec, err := unmarshal(schemaContent, "json")
	require.NoError(t, err)

	tests := []struct {
		name       string
		subject    any
		wantErrors int
		wantErr    error
	}{
		{
			name: "valid subject returns no errors",
			subject: map[string]any{
				"values": []any{
					map[string]any{"name": "input"},
				},
			},
			wantErrors: 0,
		},
		{
			name:       "missing required field values returns error",
			subject:    map[string]any{"name": "input"},
			wantErrors: 1,
		},
		{
			name: "missing required name in values item returns error",
			subject: map[string]any{
				"values": []any{
					map[string]any{"other": "field"},
				},
			},
			wantErrors: 1,
		},
		{
			name: "wrong type for values returns error",
			subject: map[string]any{
				"values": "not-an-array",
			},
			wantErrors: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := validate(spec, tt.subject)

			require.ErrorIs(t, err, tt.wantErr)
			assert.Len(t, got, tt.wantErrors)
		})
	}
}

func TestWriteErrors(t *testing.T) {
	t.Parallel()

	errors := []validationError{
		{Field: "values", Message: "is required"},
	}

	tests := []struct {
		name    string
		errors  []validationError
		format  string
		wantErr error
	}{
		{
			name:   "writes json output",
			errors: errors,
			format: "json",
		},
		{
			name:   "writes yaml output",
			errors: errors,
			format: "yaml",
		},
		{
			name:   "writes yml output",
			errors: errors,
			format: "yml",
		},
		{
			name:    "unsupported format returns error",
			errors:  errors,
			format:  "toml",
			wantErr: errFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rd, wr, err := os.Pipe()
			require.NoError(t, err)

			old := os.Stdout
			os.Stdout = wr

			writeErr := writeErrors(tt.errors, tt.format)

			require.NoError(t, wr.Close())

			os.Stdout = old

			require.NoError(t, rd.Close())

			require.ErrorIs(t, writeErr, tt.wantErr)
		})
	}
}
