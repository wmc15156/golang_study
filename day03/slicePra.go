package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int {1, 2, 3, 4, 5}
	slice2 := []int {4, 5, 6, 7, 8}
	slice3 := []int {9, 10, 11 ,12, 13}

	slice1 = append(slice1, 100,200)
	slice2 = append(slice2, slice1...)
	slice3 = append(slice3, slice2[0:5]...)

	fmt.Println(slice1, slice2, slice3)

	// 정렬sort

	slice4 := []int{5,3,4,1,10,7}
	slice5 := []string{"b", "a", "d"}

	if sort.IntsAreSorted(slice4) == false  { // 숫자가 정렬이 되었는지
		sort.Ints(slice4)
	}
	if sort.StringsAreSorted(slice5) == false { // 문자열이 정렬이 되어 있는지 확인
		sort.Strings(slice5)
	}

	fmt.Println(slice4, slice5)

	// 슬라이스 복사
	// copy (복사대상, 원본)
	// make로 공간을 할당 후 복사해야 한다.
	// 복사 된 슬라이스값을 변경해도 원본값은 변하지 않음

	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := make([]int, 3, 10)
	slice8 := []int{}

	copy(slice7, slice6) // [1,2,3] 할당된 길이만큼만 복사
	copy(slice8, slice6) // []

	fmt.Println(slice7, slice8) // [1,2,3], []

	slice7[0] = 5
	slice7[1] = 6

	fmt.Println(slice7, slice6) // [5, 6, 3] [1 2 3 4 5] //원본에는 이상 무

	slice9 := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice0 := slice9[0:5:7] // index0 ~ 4까지 슬라이싱하고 용량은 7로할당

	slice0[3] = 0

	fmt.Println(slice9, slice0, cap(slice0)) //[1 2 3 0 5 6 7 8 9] [1 2 3 0 5] 7

}
