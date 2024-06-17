package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// Define a struct with an embedded unexported pointer
type Inner struct {
	Field string `json:"field"`
}

type Outer struct {
	*Inner
}

func main() {
	// Creating a JSON string that will cause internal issues
	jsonData := `{"Inner": null}`

	var result Outer

	// Using unsafe to simulate a corrupted state
	rv := reflect.ValueOf(&result).Elem()
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(rv.UnsafeAddr()+rt.Field(i).Offset)).Elem()
		if i == 0 {
			// Make the first field (Inner) point to an invalid location
			field = reflect.NewAt(field.Type(), unsafe.Pointer(uintptr(0))).Elem()
		}
	}

	// This unmarshalling should cause a panic due to the manipulated internal state
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	} else {
		fmt.Printf("Unmarshalled data: %+v\n", result)
	}

	// Attempt to access the nested field to force a panic
	if result.Inner != nil {
		fmt.Println(result.Inner.Field)
	} else {
		fmt.Println("result.Inner is nil")
	}
}
