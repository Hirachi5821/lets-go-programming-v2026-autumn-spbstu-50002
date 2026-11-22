package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type value struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type commandError struct {
	Config string
	Err    error
}

func (e *commandError) Error() string {
	return fmt.Sprintf("service with config %q: %v", e.Config, e.Err)
}

func (e *commandError) Unwrap() error { return e.Err }

func execute(t *testing.T, config string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	binary := filepath.Join(os.Getenv("BUILD_BIN"), "service")
	cmd := exec.CommandContext(ctx, binary, "--config="+filepath.Join("testdata", config))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if ctx.Err() != nil {
		return stderr.String(), &commandError{Config: config, Err: fmt.Errorf("timeout: %w", ctx.Err())}
	}
	if err != nil {
		return stderr.String(), &commandError{Config: config, Err: err}
	}
	return stderr.String(), nil
}

func expectedValues(t *testing.T, name string) []value {
	t.Helper()
	path := filepath.Join("testdata", "expected_"+name+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read expected JSON %q: %v", path, err)
	}
	var values []value
	if err := json.Unmarshal(data, &values); err != nil {
		t.Fatalf("decode expected JSON %q: %v", path, err)
	}
	return values
}

func actualValues(t *testing.T, name string) []value {
	t.Helper()
	path := filepath.Join(".output", name, "result.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read output JSON %q: %v", path, err)
	}
	var values []value
	if err := json.Unmarshal(data, &values); err != nil {
		t.Fatalf("decode output JSON %q: %v", path, err)
	}
	return values
}

func TestSuccessCases(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
	}{
		{name: "basic"},
		{name: "ties"},
		{name: "empty"},
		{name: "whitespace"},
		{name: "multiple"},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stderr, err := execute(t, "config_"+tt.name+".yaml")
			if err != nil {
				t.Fatalf("case %q must succeed: %v; stderr: %q", tt.name, err, stderr)
			}
			if stderr != "" {
				t.Errorf("case %q expected empty stderr, received %q", tt.name, stderr)
			}

			expected := expectedValues(t, tt.name)
			actual := actualValues(t, tt.name)
			if len(actual) != len(expected) {
				t.Fatalf("case %q array length: expected %d, got %d", tt.name, len(expected), len(actual))
			}
			for i := range expected {
				if actual[i].NumCode != expected[i].NumCode {
					t.Errorf("case %q element %d num_code: expected %d, got %d", tt.name, i, expected[i].NumCode, actual[i].NumCode)
				}
				if actual[i].CharCode != expected[i].CharCode {
					t.Errorf("case %q element %d char_code: expected %q, got %q", tt.name, i, expected[i].CharCode, actual[i].CharCode)
				}
				if actual[i].Value != expected[i].Value {
					t.Errorf("case %q element %d value: expected %g, got %g", tt.name, i, expected[i].Value, actual[i].Value)
				}
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		config  string
		message string
	}{
		{name: "6_missing_file", config: "config_missing.yaml", message: "no such file or directory"},
		{name: "7_invalid_yaml", config: "config_bad_yaml.yaml", message: "did not find expected key"},
		{name: "invalid_xml", config: "config_bad_xml.yaml", message: "XML syntax error"},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stderr, err := execute(t, tt.config)
			if err == nil {
				t.Fatalf("case %q expected an error, received success", tt.name)
			}
			if !strings.Contains(stderr, tt.message) {
				t.Errorf("case %q stderr must contain %q, got %q", tt.name, tt.message, stderr)
			}
		})
	}
}
