package json

import (
	JSON "encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/Jeffail/gabs/v2"
)

// Build JSON from map of keys & translations
func Build(data map[string]string) ([]byte, error) {
	return JSON.Marshal(data)
}

// Build nested JSON from map of keys & translations
func BuildNested(data map[string]string) ([]byte, error) {
	result := gabs.New()

	for key, value := range data {
		result.SetP(value, key)
	}

	return result.MarshalJSON()
}

type RawJSON map[string]interface{}

// Build map of keys & translations from JSON
func Parse(data []byte) (*map[string]string, error) {
	var rawJSON RawJSON

	err := JSON.Unmarshal(data, &rawJSON)

	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for k, raw := range rawJSON {
		flatten(result, k, reflect.ValueOf(raw))
	}

	return &result, nil
}

func flatten(result map[string]string, prefix string, v reflect.Value) error {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			result[prefix] = "true"
		} else {
			result[prefix] = "false"
		}
	case reflect.Int:
		result[prefix] = fmt.Sprintf("%d", v.Int())
	case reflect.Float64:
		result[prefix] = fmt.Sprintf("%f", v.Float())
	case reflect.Map:
		err := flattenMap(result, prefix, v)
		if err != nil {
			return err
		}
	case reflect.Slice:
		err := flattenSlice(result, prefix, v)
		if err != nil {
			return err
		}
	case reflect.String:
		result[prefix] = v.String()
	default:
		return errors.New("Could not process some data")
	}

	return nil
}

func flattenMap(result map[string]string, prefix string, v reflect.Value) error {
	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}

		if k.Kind() != reflect.String {
			return errors.New("Could not process some data")
		}

		err := flatten(result, fmt.Sprintf("%s.%s", prefix, k.String()), v.MapIndex(k))

		if err != nil {
			return err
		}
	}

	return nil
}

func flattenSlice(result map[string]string, prefix string, v reflect.Value) error {
	prefix = prefix + "."

	for i := 0; i < v.Len(); i++ {
		err := flatten(result, fmt.Sprintf("%s%d", prefix, i), v.Index(i))

		if err != nil {
			return err
		}
	}

	return nil
}
