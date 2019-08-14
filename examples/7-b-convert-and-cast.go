package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// conversion
	// occurs when converting a value into another type

	myString := "hello world!"
	fmt.Println("myString", myString)

	// convert "string" to "[]byte"
	asByteArray := []byte(myString)
	fmt.Println("asByteArray", asByteArray)

	// convert "[]byte" to "string"
	asString := string(asByteArray)
	fmt.Println("asString", asString)

	// -----------------------

	// casting
	// occurs when reading data out

	// setup
	asByte := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var asMap map[string]interface{}
	if err := json.Unmarshal(asByte, &asMap); err != nil {
		panic(err)
	}
	fmt.Println(asMap)

	// as "float64"
	asFloat64 := asMap["num"].(float64)
	fmt.Println(asFloat64)

	// this will fail
	// "num" is not a string
	asInt := asMap["num"].(string)
	fmt.Println(asInt)
}
