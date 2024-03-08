package config

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func ParseTag(tag string, conf any) error {
	_, err := requiredFields(tag, reflect.ValueOf(conf))
	if err != nil {
		return err
	}
	return nil
}

// required fileds can not be zero.
// if pass, this function will return all required fields.
func requiredFields(tag string, conf reflect.Value) ([]string, error) {
	resultFields := make([]string, 0)

	// Use dereferenced type in case of pointer receivers.
	rt := conf.Type()
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
		conf = conf.Elem()
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		// Dereference field value when it's a pointer.
		fieldValue := conf.Field(i)
		if fieldValue.Kind() == reflect.Pointer && !fieldValue.IsNil() {
			fieldValue = fieldValue.Elem()
		}

		tagValue := field.Tag.Get(tag)
		if tagValue != "" {
			values := strings.Split(tagValue, ",")
			if slices.Contains(values, "required") {
				if fieldValue.IsZero() {
					return nil, fmt.Errorf("field %s is required", field.Name)
				}
				resultFields = append(resultFields, field.Name)
			}
		}

		// Check recursively if the field is a nested struct and not a time.Time (special case).
		if field.Type.Kind() == reflect.Struct && field.Type.String() != "time.Time" {
			nestedRequiredFields, err := requiredFields(tag, fieldValue)
			if err != nil {
				return nil, err
			}
			for _, nestedField := range nestedRequiredFields {
				resultFields = append(resultFields, field.Name+"."+nestedField)
			}
		}
	}
	return resultFields, nil
}
