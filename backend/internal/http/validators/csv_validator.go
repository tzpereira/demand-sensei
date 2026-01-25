package validators

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"reflect"
	"strings"

	models "demand-sensei/backend/internal/http/schemas"
)

// Returns a generic CSV validator that resolves schema by import type
func GetValidator(importType string) func(*multipart.FileHeader) error {
	schema, err := getSchemaByImportType(importType)
	if err != nil {
		return func(file *multipart.FileHeader) error {
			return err
		}
	}

	return func(file *multipart.FileHeader) error {
		return validateCSVHeaders(file, schema)
	}
}

// Maps import type -> schema
func getSchemaByImportType(importType string) (interface{}, error) {
	switch importType {
	case "sales":
		return models.SalesImportSchema{}, nil
	default:
		return nil, errors.New("invalid import type")
	}
}

// Validates CSV headers against required fields defined in struct tags
func validateCSVHeaders(file *multipart.FileHeader, schema interface{}) error {
	f, err := file.Open()
	if err != nil {
		return errors.New("failed to open file: " + err.Error())
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	headers, err := reader.Read()
	if err != nil {
		return errors.New("failed to read CSV header: " + err.Error())
	}

	normalizedHeaders := normalize(headers)
	requiredFields := getRequiredCSVFields(schema)

	var missing []string
	for _, field := range requiredFields {
		if !contains(normalizedHeaders, field) {
			missing = append(missing, field)
		}
	}

	if len(missing) > 0 {
		return errors.New("missing required fields: " + strings.Join(missing, ", "))
	}

	return nil
}

// Extracts required CSV fields from struct tags
func getRequiredCSVFields(schema interface{}) []string {
	t := reflect.TypeOf(schema)
	fields := []string{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if strings.Contains(f.Tag.Get("validate"), "required") {
			csvTag := strings.Split(f.Tag.Get("csv"), ",")[0]
			if csvTag != "" {
				fields = append(fields, normalizeValue(csvTag))
			}
		}
	}

	return fields
}

// Helpers
func normalize(values []string) []string {
	out := make([]string, 0, len(values))
	for _, v := range values {
		out = append(out, normalizeValue(v))
	}
	return out
}

func normalizeValue(v string) string {
	return strings.ToLower(strings.TrimSpace(v))
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
