package handlers

import (
	"reflect"
	"strings"
)

func in(where, what interface{}) bool {
	// t := reflect.ValueOf(where)
	// switch t.Kind() {
	switch where.(type) {
	case *string:
		return strings.Contains(where.(string), what.(string))
	case *[]interface{}:
		for _, v := range where.([]interface{}) {
			if v == what {
				return true
			}
		}
		return false
	case *map[interface{}]interface{}:

	// 	key_type := reflect.TypeOf(where).Key()
	// 	value_type := reflect.TypeOf(where).Elem()
	// 	if key_type.Kind() == reflect.String && value_type.Kind() == reflect.Float64 {
	// 		for _, key := range where.(map[string]float64) {
	// 			if key == what {
	// 				return true
	// 			}
	// 		}
	// 	} else if key_type.Kind() == reflect.String && value_type.Kind() == reflect.String {
	// 		for _, key := range where.(map[string]string) {
	// 			if key == what {
	// 				return true
	// 			}
	// 		}
	// 	} else if key_type.Kind() == reflect.Int64 && value_type.Kind() == reflect.Float64 {
	// 		for _, key := range where.(map[int64]float64) {
	// 			if key == what {
	// 				return true
	// 			}
	// 		}
	// 	} else {
	// 		return false
	// 	}
	// }
		return false
	return false
}
