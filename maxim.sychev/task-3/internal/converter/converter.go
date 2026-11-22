package converter

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type FileError struct {
	Operation string
	Path      string
	Err       error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s %q: %v", e.Operation, e.Path, e.Err)
}

func (e *FileError) Unwrap() error { return e.Err }

type DecodeError struct {
	Format string
	Err    error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("decode %s: %v", e.Format, e.Err)
}

func (e *DecodeError) Unwrap() error { return e.Err }

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type currenciesXML struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Items   []currencyXML `xml:"Valute"`
}

type currencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Run(configPath string) error {
	if configPath == "" {
		return &DecodeError{Format: "configuration", Err: fmt.Errorf("--config is required")}
	}

	cfg, err := readConfig(configPath)
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	currencies, err := readCurrencies(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("load currencies: %w", err)
	}

	if err := writeCurrencies(cfg.OutputFile, currencies); err != nil {
		return fmt.Errorf("save currencies: %w", err)
	}

	return nil
}

func readConfig(path string) (config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, &FileError{Operation: "read", Path: path, Err: err}
	}

	var cfg config
	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	decoder.KnownFields(true)
	if err := decoder.Decode(&cfg); err != nil {
		return config{}, &DecodeError{Format: "YAML", Err: err}
	}
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return config{}, &DecodeError{Format: "YAML", Err: fmt.Errorf("input-file and output-file are required")}
	}

	return cfg, nil
}

func readCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, &FileError{Operation: "read", Path: path, Err: err}
	}

	var input currenciesXML
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&input); err != nil {
		return nil, &DecodeError{Format: "XML", Err: err}
	}

	result := make([]Currency, 0, len(input.Items))
	for _, item := range input.Items {
		currency, err := parseCurrency(item)
		if err != nil {
			return nil, fmt.Errorf("parse currency %q: %w", item.CharCode, err)
		}
		result = append(result, currency)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Value == result[j].Value {
			return result[i].CharCode < result[j].CharCode
		}
		return result[i].Value > result[j].Value
	})
	return result, nil
}

func parseCurrency(item currencyXML) (Currency, error) {
	numCode, err := strconv.Atoi(strings.TrimSpace(item.NumCode))
	if err != nil {
		return Currency{}, &DecodeError{Format: "NumCode", Err: err}
	}

	valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return Currency{}, &DecodeError{Format: "Value", Err: err}
	}

	return Currency{NumCode: numCode, CharCode: strings.TrimSpace(item.CharCode), Value: value}, nil
}

func writeCurrencies(path string, currencies []Currency) error {
	parent := filepath.Dir(path)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return &FileError{Operation: "create directory", Path: parent, Err: err}
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return &DecodeError{Format: "JSON", Err: err}
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return &FileError{Operation: "write", Path: path, Err: err}
	}
	return nil
}
