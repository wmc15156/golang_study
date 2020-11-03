package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	name string
	balance float64
	interest float64
}

type car struct {
	color string "차랭색"
	name string  "차량이름"
	company string "차량회사"
	details spec "차량스펙"
}
type spec struct {
	length int
	height int
	width int
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 예제1
	kim := Account{name: "1234-123", balance: 10000, interest: 0.15}
	lee := Account{name: "456-798", balance: 2000, interest: 0.2}

	fmt.Println(kim, lee)

	fmt.Println(kim.Calculate())
	fmt.Println(lee.Calculate())

	// 예제2

	var park *Account = new(Account)
	park.name = "1234-123"
	park.balance = 10000
	park.interest = 0.01

	hong := &Account{name: "456-798", balance: 4000, interest: 0.4}

	james := new(Account)
	james.name = "44432-1234"
	james.balance = 2000
	james.interest = 0.2

	fmt.Println(park)
	fmt.Println(hong)
	fmt.Println(james)
	fmt.Printf("%#v\n", park)
	fmt.Printf("%#v\n", hong)
	fmt.Printf("%#v\n", james)

	fmt.Println()

	// 예제3

	c1 := car{color: "red", name: "520d"}
	c2 := new(car)
	c2.color, c2.name = "white", "530d"
	c3 := &car{}
	c4 := &car{color: "green", name: "mint"}

	fmt.Println(c1, c2, c3, c4)

	fmt.Printf("%#v\n", c2)

	// 예제4
	cars := []struct{name string; color string}{{"520d", "red"},{"530i", "blue"}, {"540d", "white"}}

	for _, value := range(cars) {
		fmt.Println(value.name)
		fmt.Printf("%#v", value)
		fmt.Println()
	}

 	// 예제5

 	car2 := car{"green", "520d", "bmw", spec{length: 1000, height: 1000, width: 1000}}
	tag := reflect.TypeOf(car2)

 	for i:=0; i < tag.NumField(); i++ {
 		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

	fmt.Println(car2)
 	fmt.Printf("%#v", car2)

 	fmt.Println()

 	fmt.Println(car2.details.length, car2.details.width, car2.details.height)


}