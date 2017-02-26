package analyzer

import (
	"fmt"
	"math"
)

func analyzeData(data map[string]interface{}) {
	for k, v := range data {
		switch v.(type) {
		case int, int8, int16, int32, int64:
			fmt.Println(k, v, "is int")
		case float32, float64:
			if v == math.Trunc(v.(float64)) {
				fmt.Println(k, v, "is int")
			} else {
				fmt.Println(k, v, "is float64")
			}
		case string:
			fmt.Println(k, v, "is string")
		case bool:
			fmt.Println(k, v, "is bool")
		case interface{}:
			analyzeData(v.(map[string]interface{}))
		default:
			fmt.Println(k, v, "is unknown")
		}
	}
}

// Analyze data and returns structure with the corresponding types
func Analyze(data map[string]interface{}) {
	analyzeData(data)
}
