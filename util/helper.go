package util

import (
	"os"
	"reflect"
)

// Merge two maps, overriding existing keys from b --> a into a new map
func MergeMaps(a map[string]interface{}, b map[string]interface{}) map[string]interface{} {
	retVal := make(map[string]interface{}, len(a)+len(b))
	for k, v := range a {
		retVal[k] = v
	}
	for k, v := range b {
		retVal[k] = v
	}
	return retVal
}

// Check if a given value exists in the given array
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// Iterate over a given map recursively and call the given func on every value
func WalkRecursive(values map[string]interface{}, fn func(interface{})) {
	for _, v := range values {
		switch v.(type) {
		case map[string]interface{}:
			WalkRecursive(v.(map[string]interface{}), fn)
		default:
			fn(&v)
		}
	}
}

// Get an environment variable or return the default value
func GetEnvOrDefault(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
