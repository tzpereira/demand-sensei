package validators

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"reflect"
	"strings"

	"demand-sensei/backend/internal/http/models"
)

func GetValidator(importType string) func(*multipart.FileHeader) error {
	switch importType {
	case "sales":
		return ValidateSales
	default:
		return func(file *multipart.FileHeader) error {
			return errors.New("invalid import type")
		}
	}
}

func ValidateSales(file *multipart.FileHeader) error {
	return validateCSV(file, models.SalesRecord{})
}

func validateCSV(file *multipart.FileHeader, model interface{}) error {
	f, err := file.Open()
	if err != nil {
		return errors.New("failed to open file: " + err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	headers, err := r.Read()
	if err != nil {
		return errors.New("failed to read CSV header: " + err.Error())
	}

	requiredFields := getRequiredCSVFields(model)

	missing := []string{}
	for _, field := range requiredFields {
		if !contains(headers, field) {
			missing = append(missing, field)
		}
	}

	if len(missing) > 0 {
		return errors.New("missing required fields: " + strings.Join(missing, ", "))
	}

	return nil
}

func getRequiredCSVFields(model interface{}) []string {
	t := reflect.TypeOf(model)
	fields := []string{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if strings.Contains(f.Tag.Get("validate"), "required") {
			csvTag := strings.Split(f.Tag.Get("csv"), ",")[0]
			if csvTag != "" {
				fields = append(fields, csvTag)
			}
		}
	}

	return fields
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
