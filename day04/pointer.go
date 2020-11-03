package main

import "fmt"

func main() {
	// 포인터
	// 변수의 지역성, 연속된 메모리 참조// 다른 함수에서 사용(?)
	// 주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그방지)
	// *(애스터리스크로)사용
	// nil로 초기화

	i := 7
	a := &i
	b := &i
	c := &i
	d := &i

	fmt.Println(a, b, c, d, &i) // 0xc0000180b8 0xc0000180b8 0xc0000180b8 0xc0000180b8 0xc0000180b8

	fmt.Println(*a, *b, *c, *d) //역참조 7 7 7 7

	*d = 10

	fmt.Println(*a, *b, *c, *d, i) // 10 10 10 10 10




}
