package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type User struct {
	Name    string `json:"name" required:"Name is required custom"`
	Email   string `json:"email" required:"Email is required custom"`
	Profile string `json:"profile" required:"Profile is required custom"`
	Id      int    `json:"id" required:"id is required"`
}

func ValidateJSON[T any](payload io.Reader) error {
	var body T

	decoder := json.NewDecoder(payload)
	err := decoder.Decode(&body)
	if err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			return fmt.Errorf(
				"field '%s' has an invalid type: expected %s but got %s",
				e.Field, e.Type.String(), e.Value,
			)
		case *json.SyntaxError:
			return fmt.Errorf("syntax error at offset %d: %v", e.Offset, err)
		case *json.InvalidUnmarshalError:
			return fmt.Errorf("invalid unmarshal: %v", e)
		case *json.UnsupportedTypeError:
			return fmt.Errorf("unsupported type: %v", e.Type)
		case *json.MarshalerError:
			return fmt.Errorf("error marshaling JSON: %v", e.Err)
		default:
			return fmt.Errorf("unexpected error: %v", err)
		}
	}

	v := reflect.ValueOf(body)
	t := reflect.TypeOf(body)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("json")
		required := field.Tag.Get("required")
		if required != "" {
			if isEmptyValue(value) {
				return fmt.Errorf("field '%s' (%s) is required but missing", field.Name, tag)
			}
		}
		switch value.Kind() {
		case reflect.String:
			if required == "true" && value.String() == "" {
				return fmt.Errorf("field '%s' (%s) must be a non-empty string", field.Name, tag)
			}
		case reflect.Int, reflect.Int64:
			if required == "true" && value.Int() == 0 {
				return fmt.Errorf("field '%s' (%s) must be a non-zero integer", field.Name, tag)
			}
		default:
			fmt.Printf("Field '%s' is of type %s\n", field.Name, value.Kind())
		}
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}

func RunTestValidation() {
	customBody := `{
		"name": "John Doe",
		"email": "john.doe@example.com",
		"profile": "This",
		"id": "asjhdh"
	}`

	reader := strings.NewReader(customBody)

	err := ValidateJSON[User](reader)
	if err != nil {
		fmt.Println(err)
	}
}
