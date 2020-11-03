package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age int
	address string
	color []string
}

func main() {
	go nameCheck("Kim")
	go nameCheck("Girl")
	result := practiceFunc("skdmaksdm")
	fmt.Println(result)
	a := practiceFunc2("dd", 12,"seiul")
	fmt.Println(a)

}

func nameCheck (name string) {
	for i:=0; i<10; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func practiceFunc(name string) (result string) {
	result = name
	return
}

func practiceFunc2 (name string, age int, address string) (result []person) {
	arr := []string {"red", "green"}
	total := person {
		name: name,
		age: age,
		address: address,
		color: arr,
	}
	fmt.Println(total)
	result = append(result, total)
	return
}