package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// templateData includes our klient converts the given raw interface to a
// []byte data that can used to pass into packer.Template().
func TemplateData(raw, provisioner interface{}) ([]byte, error) {
	rawMapData, err := ToMap(raw, "packer")
	if err != nil {
		return nil, err
	}

	packerTemplate := map[string]interface{}{}
	packerTemplate["builders"] = []interface{}{rawMapData}
	packerTemplate["provisioners"] = provisioner

	return json.Marshal(packerTemplate)
}

// toMap converts a struct defined by `in` to a map[string]interface{}. It only
// extract data that is defined by the given tag.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only struct is allowd got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out, nil

}

// toUint tries to convert the given to uint type
func ToUint(x interface{}) uint {
	switch i := x.(type) {
	case float64:
		return uint(i)
	case uint:
		return i
	case int:
		return uint(i)
	case string:
		s, err := strconv.Atoi(i)
		if err != nil {
			log.Println("cannot convert", i)
			return 0
		}

		return uint(s)
	case int64:
		return uint(i)
	default:
		return 0
	}
}

func RandString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
