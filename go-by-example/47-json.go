package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// dealing with json is important
// go supports json encoding and decoding
// for builtin and custom types
// only EXPORTED fields are encoded for structs

// "interface" is an arbitrary data type
// like "any" in typescript

// simple struct
type response1 struct {
	Page   int
	Fruits []string
}

// annotated struct
// custom json member names
// IMPORTANT: the string `json:"page"` does NOT contain a space
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// encoding / stringifying
	// ---------------------------------------------------
	// "json.Marshal(goObject)" transforms to []byte
	// then "string(jsonObject)" converts to a string

	// bool
	j1, _ := json.Marshal(true)
	// fmt.Println(j1)
	s1 := string(j1)
	fmt.Println(s1)

	// int
	j2, _ := json.Marshal(1)
	// fmt.Println(j2)
	s2 := string(j2)
	fmt.Println(s2)

	// float
	j3, _ := json.Marshal(2.34)
	// fmt.Println(j3)
	s3 := string(j3)
	fmt.Println(s3)

	// string
	j4, _ := json.Marshal("gopher")
	// fmt.Println(j4)
	s4 := string(j4)
	fmt.Println(s4)

	// []string
	j5, _ := json.Marshal([]string{"apple", "peach", "pear"})
	// fmt.Println(j5)
	s5 := string(j5)
	fmt.Println(s5)

	// map[string]int
	j6, _ := json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	// fmt.Println(j6)
	s6 := string(j6)
	fmt.Println(s6)

	// custom (basic)
	// use the point like a "new"
	j7Data := &response1{
		Page:   1,
		Fruits: []string{"apple", "rambouton"}}
	j7, _ := json.Marshal(j7Data)
	// fmt.Println(j7)
	s7 := string(j7)
	fmt.Println(s7)

	// custom (annotated)
	// pointer usage again - in the place of "new"
	j8Data := &response2{
		Page:   1,
		Fruits: []string{"apple", "lychee", "longan"}}
	j8, _ := json.Marshal(j8Data)
	// fmt.Println(j8)
	s8 := string(j8)
	fmt.Println(s8)

	// decoding / jsonifying
	// ---------------------------------------------------
	// "json.Unmarshal(goObject)" transforms to []byte
	// then "string(jsonObject)" converts to a string

	// byte
	// generic data structure
	// this is how to build a map of json objects keyed by string
	// "interface{}" is a generic data structure
	// also remember to pass in the pointer for the Unmarshal
	// it will attempt to store the value in "&dat"
	// but return "err != nil" if error
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	var dat map[string]interface{}

	// do the conversion and panic if error
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// read from the map
	// convert to float64
	// conversion required b/c "interface{}" is too generic
	num := dat["num"].(float64)
	fmt.Println(num)

	// decoding nested data requires multiple conversions
	// "strs" is the decoded array object
	// "but "str1" is the first string in the array
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// custom structs
	// eliminates the need for the type-safety
	// also eliminates the need for additional conversion
	// vastly superior
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// -------------------------------------------

	// can stream json encodings
	// to os.Writers, os.Stdout
	// OR EVEN Http response bodies
	// this example will write the encoding to the terminal's stdout
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettucey": 7}
	enc.Encode(d)

}
