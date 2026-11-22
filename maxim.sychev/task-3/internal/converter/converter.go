package converter

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	directoryPermissions = 0o755
	outputPermissions    = 0o600
)

var (
	errConfigRequired = errors.New("--config is required")
	errConfigFields   = errors.New("input-file and output-file are required")
)

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
		return errConfigRequired
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
		return config{}, fmt.Errorf("read configuration %q: %w", path, err)
	}

	var cfg config

	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	decoder.KnownFields(true)

	if err := decoder.Decode(&cfg); err != nil {
		return config{}, fmt.Errorf("decode YAML: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return config{}, errConfigFields
	}

	return cfg, nil
}

func readCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	var input currenciesXML

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&input); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
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
	numCode := 0

	if code := strings.TrimSpace(item.NumCode); code != "" {
		parsed, err := strconv.Atoi(code)
		if err != nil {
			return Currency{}, fmt.Errorf("invalid NumCode %q: %w", item.NumCode, err)
		}

		numCode = parsed
	}

	valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return Currency{}, fmt.Errorf("invalid Value %q: %w", item.Value, err)
	}

	return Currency{NumCode: numCode, CharCode: strings.TrimSpace(item.CharCode), Value: value}, nil
}

func writeCurrencies(path string, currencies []Currency) error {
	parent := filepath.Dir(path)
	if err := os.MkdirAll(parent, directoryPermissions); err != nil {
		return fmt.Errorf("create output directory %q: %w", parent, err)
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, outputPermissions); err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}
