package flattenjson

import (
	"encoding/json"
	"errors"
	"strconv"
)

func flattenList(jsonSlice []interface{}, prefix string, delimiter string, flattenSlice bool) (map[string]interface{}, error) {
	var err error
	var sliceKey string
	finalMap := make(map[string]interface{})
	for key, value := range jsonSlice {
		if len(prefix) > 0 {
			sliceKey = prefix + delimiter + strconv.Itoa(key)
		} else {
			sliceKey = strconv.Itoa(key)
		}
		switch v := value.(type) {
		case nil:
			finalMap[sliceKey] = v
		case int:
			finalMap[sliceKey] = v
		case float64:
			finalMap[sliceKey] = v
		case string:
			finalMap[sliceKey] = v
		case bool:
			finalMap[sliceKey] = v
		case []interface{}:
			if flattenSlice {
				out := make(map[string]interface{})
				out, err = flattenList(v, sliceKey, delimiter, flattenSlice)
				if err != nil {
					return nil, err
				}
				for keyInner, valueInner := range out {
					finalMap[keyInner] = valueInner
				}
			} else {
				finalMap[sliceKey] = v
			}

		case map[string]interface{}:
			out := make(map[string]interface{})
			out, err = flattenMap(v, sliceKey, delimiter, flattenSlice)
			if err != nil {
				return nil, err
			}
			for keyInner, valueInner := range out {
				finalMap[keyInner] = valueInner
			}
		}
	}
	return finalMap, nil
}

func flattenMap(jsonMap map[string]interface{}, prefix string, delimiter string, flattenSlice bool) (map[string]interface{}, error) {
	var err error
	finalMap := make(map[string]interface{})
	for key, value := range jsonMap {
		if len(prefix) > 0 {
			key = prefix + delimiter + key
		}
		switch v := value.(type) {
		case nil:
			finalMap[key] = v
		case int:
			finalMap[key] = v
		case float64:
			finalMap[key] = v
		case string:
			finalMap[key] = v
		case bool:
			finalMap[key] = v
		case []interface{}:
			if flattenSlice {
				out := make(map[string]interface{})
				out, err = flattenList(v, key, delimiter, flattenSlice)
				if err != nil {
					return nil, err
				}
				for keyInner, valueInner := range out {
					finalMap[keyInner] = valueInner
				}
			} else {
				finalMap[key] = v
			}

		case map[string]interface{}:
			out := make(map[string]interface{})
			out, err = flattenMap(v, key, delimiter, flattenSlice)
			if err != nil {
				return nil, err
			}
			for keyInner, valueInner := range out {
				finalMap[keyInner] = valueInner
			}
		}
	}
	return finalMap, nil
}

// JSONByte takes in a  nested JSON as a byte array and a delimiter and returns an
// flattened json byte array
// jsonByte = JSON byte array
// delimeter = character used to separate keys
// flattenSlice = bool to decide weather to flatten slices within the JSON
func JSONByte(jsonByte []byte, delimeter string, flattenSlice bool) ([]byte, error) {
	var input interface{}
	var m map[string]interface{}
	var out []byte

	err := json.Unmarshal(jsonByte, &input)
	if err != nil {
		return nil, err
	}
	switch t := input.(type) {
	case map[string]interface{}:
		m, err = flattenMap(t, "", delimeter, flattenSlice)
		if err != nil {
			return nil, err
		}
	case []interface{}:
		m, err = flattenList(t, "", delimeter, flattenSlice)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid JSON")
	}
	out, err = json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return out, nil

}
