package main

import "fmt"


func sum(nums ...int) int {
	fmt.Println(nums, "nums")
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func mulitValue(i *int) {
	*i *= 10
}

func multiplyValue(a ,b int) (r1 int, r2 int){
	r1 = a * 10
	r2 = b * 20
	return
}

func prtWord(msg ...string) {
	fmt.Println(msg)
	for _, value := range msg {
		fmt.Print(" ", value)
	}
}

func sum1(x, y int) (r int) {
	r = x + y
	return
}

func multiply1(x, y int) (r int) {
	r = x * y
	return
}

func main() {

	c , d := multiplyValue(10, 20)
	fmt.Println(c, d ,"마지막값")

	// 가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
	a := sum(1, 2, 3, 4, 5)
	fmt.Println(a)

	prtWord("a", "apple", "banana", "test", "golang")

	f := []int{1, 2, 3, 4, 5}

	b := sum(f...)
	fmt.Println(b)

	fu := []func(int, int) int{multiply1, sum1}


	test := fu[0](20, 20)
	test2 := fu[1](10, 10)

	fmt.Println("test1,2", test, test2)

	mapTest := map[string]func(int, int ) int {
		"multy": multiply1,
		"sum": sum1,
	}

	fmt.Println(mapTest["multy"](10, 10))

}


