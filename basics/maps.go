<<<<<<< HEAD
package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11

	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, unknowvalue := myMap["key1"]

	if unknowvalue {
		fmt.Println("A value exist with key1")
	} else {
		fmt.Println("No value exist with key1")
	}

	// fmt.Println(value)
	fmt.Println("true/false", unknowvalue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is'n initialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val)
}
=======
package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11

	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, unknowvalue := myMap["key1"]

	if unknowvalue {
		fmt.Println("A value exist with key1")
	} else {
		fmt.Println("No value exist with key1")
	}

	// fmt.Println(value)
	fmt.Println("true/false", unknowvalue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is'n initialized to nil value")
	}

	val := myMap4["key"]
	fmt.Println(val)
}
>>>>>>> fb36f3b (wls_push)
