package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// encode
	fmt.Printf("---encode---\n")

	// bool
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// int
	intB, _ := json.Marshal(10)
	fmt.Println(string(intB))
	
	// float
	floB, _ := json.Marshal(2.14)
	fmt.Println(string(floB))

	// string
	strB, _ := json.Marshal("Hi!")
	fmt.Println(string(strB))

	// slice
	slc := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slc)
	fmt.Println(string(slcB))

	// map
	m := map[string]int{"apple": 5, "lettuce": 1}
	mapB, _ := json.Marshal(m)
	fmt.Println(string(mapB))

	// custom type
	res1 := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1)
	fmt.Println(string(res1B))

	res2 := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2)
	fmt.Println(string(res2B))


	// decode
	fmt.Printf("\n\n---decode---\n")

	byts := []byte(`{"num": 3.14, "strs": ["apple", "beach", "pear]"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byts, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// assert number
	num, ok := dat["num"].(float64)
	if !ok {
		fmt.Println("num is not float64")
	} else {
		fmt.Println(num)
	}

	// assert []string
	strs, ok := dat["strs"].([]interface{})
	if !ok {
		fmt.Println("strs is not []interface{}")
	} else {
		fmt.Println(strs)
		
		// assert element of []string
		str, ok := strs[0].(string)
		if !ok {
			fmt.Println("element of []interface{} is not string")
		} else {
			fmt.Println(str)
		}
	}

	// decode to custom type
	str := `{"page": 1, "fruits": ["apple", "peach", "pear"]}`
	var res response1
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Printf("res.Page = %d, res.Fruits = %s\n", res.Page, res.Fruits)

	// use Encoder to encode JSON string
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 1, "peach": 3}
	if err := enc.Encode(d); err != nil {
		panic(err)
	}
}
