package myFunctions

import (
	"fmt"
)

func SwitchAllTypr(iV interface{}) {
	// 这里的 v 是类型的值
	switch nameiV := iV.(type) {
	case bool:
		fmt.Printf("bool\n")
	case string:
		fmt.Printf("string\n")
	case int, int64:
		fmt.Printf("int or int 64\n")
	case float32, float64:
		fmt.Printf("float32 or 64\n")
	default:
		fmt.Printf("what Unknown is %v\n", nameiV)
	}
}
