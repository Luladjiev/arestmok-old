package analyzer

import "math"
import "fmt"
import "strconv"

func getType(data interface{}) string {
	switch data.(type) {
	case int, int8, int16, int32, int64:
		return "int"
	case float32, float64:
		if data == math.Trunc(data.(float64)) {
			return "int"
		}
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case []interface{}:
		return "array"
	case interface{}:
		return "object"
	default:
		return "unknown"
	}
}

func analyzeData(data map[string]interface{}, property string) {
	for k, v := range data {
		if property != "" {
			k = property + "." + k
		}

		t := getType(v)

		fmt.Println(k, v, t)

		switch t {
		case "array":
			for i, e := range v.([]interface{}) {
				elementType := getType(e)
				fmt.Println(k, i, e, elementType)
				if elementType == "array" || elementType == "object" {
					analyzeData(e.(map[string]interface{}), k+"["+strconv.Itoa(i)+"]")
				}
			}
		case "object":
			analyzeData(v.(map[string]interface{}), k)
		}
	}
}

// Analyze data and returns structure with the corresponding types
func Analyze(data map[string]interface{}) {
	analyzeData(data, "")
}
