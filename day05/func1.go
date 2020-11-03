package main

import (
	"fmt"
)

func main() {
	// 함수선언 방법
	// 선언: func 키워드로 선언
	// 타 언어와 달리 반환 값 다수 가능
	helloGoLang()
	say_one("string")

	result := sum(5, 5)
	fmt.Println("result: ", result)
	// 숫자를 문자열로 변환

}

// 리턴값도 없고 매개변수도 없을 경우

func helloGoLang() {
	fmt.Println("Hello world")
}


func say_one(m string) {
	fmt.Println(m, "m")
}

func sum(x int, y int) int {
	return x + y
}

