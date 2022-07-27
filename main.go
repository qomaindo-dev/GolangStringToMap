package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	elementMap := make(map[string]interface{})
	dataString := "Andi, Budi, Tini, Tuti"
	arr := strings.Split(dataString, ",")
	elementMap["prefix"] = arr
	fmt.Println("Prefix", elementMap)
	fmt.Println(arr)
	fmt.Println("Data type dataString => ", reflect.TypeOf(dataString))
	fmt.Println("Data type elementMap => ", reflect.TypeOf(elementMap))
}
