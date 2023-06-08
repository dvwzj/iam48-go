package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	fields := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	first := true
	for i := 0; i < fields.NumField(); i++ {
		if !first {
			buffer.WriteString(",")
		}
		first = false
		field := fields.Field(i)
		value := values.Field(i)
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":", key))
		if field.Type.Kind() == reflect.Ptr {
			if value.IsNil() {
				elementKind := field.Type.Elem().Kind()
				if elementKind == reflect.Slice {
					buffer.WriteString("[]")
				} else if elementKind == reflect.Map || elementKind == reflect.Struct {
					buffer.WriteString("{}")
				} else {
					buffer.WriteString("null")
				}
				continue
			}
		}
		b, err := json.Marshal(value.Interface())
		if err != nil {
			return nil, err
		}
		buffer.Write(b)
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
