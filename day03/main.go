package main

import "fmt"

func main() {
	// 첫번째 방법
	var slice []string
	slice2 := []int{1,2,3,4,5} // slice

	slice3 := [][]int {
		{1,2,3,4,5},
		{4,5,6,7,8},
	}
	// slice4 := []int{} // panic: runtime error: index out of range [0] with length 0
	// slice4[0] = 1 // error발생 길이가 가변형이므로 원소를 넣지 않는 상태에서 할당하면 에러발생

	var arr1 = [5]int{1,2,3}


	fmt.Println(slice, len(slice), cap(slice), slice)
	fmt.Println(slice2, len(slice2), cap(slice2), slice2)
	fmt.Println(slice3, len(slice3), cap(slice3), slice3)
	fmt.Println(arr1, len(arr1), cap(arr1), arr1)

	// 예제2

	var slice5 []int = make([]int, 5, 10) // 길이는 5 용량은 10을 확보(생략시 길이와 같게 저장)
	var slice6 = make([]int, 5,10)
	slice7 := make([]int,5, 100)

	fmt.Println(slice5, len(slice5), cap(slice5), slice5)
	fmt.Println(slice6, len(slice6), cap(slice6), slice6)
	fmt.Println(slice7, len(slice7), cap(slice7), slice7)

	// 예제3

	var slice8 []int
	var slice9[]int = []int{1,2,3,4,5}
	fmt.Println(slice8, slice9)


	array1 := [5]int{1,2,3,4,5}

	array2 := [5]int{}

	array2 = array1

	array2[0] = 9

	fmt.Println(array1, array2) // [1 2 3 4 5] [9 2 3 4 5]
	// ======================== //
	Slice := []int{1,2,3,4,5}
	Slice2 := []int{}

	Slice2 = Slice
	Slice2[0] = 9

	fmt.Println(Slice, Slice2) //[9 2 3 4 5] [9 2 3 4 5]

	// Slice3 := make([]int, 5, 10)
	// Slice3[6] = 5 // index out of range [6] with length 5 메모리는 용량은 10을 할당받았지만 길이는 5라서 에러발생

	// slice순회

	for i, v := range Slice {
		fmt.Println("index: ",i, "value:",v)
	}







}
