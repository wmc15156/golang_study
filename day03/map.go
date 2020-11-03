package main

import "fmt"

func main() {
	// map은 key와 value형으로 이루어져있음
	// 맵: 파이썬에서 딕셔너리와 비슷함.
	// Key-Value로 자료저장 // 참조타입
	// 순서없음

	var map1 map[string] int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	map3["a"] = 2
	fmt.Println(map1, map2, map3)

	map4 := map[string]int{}
	map4["age"] = 2
	map4["phone"] = 1234

	map5 := map[string]int {
		"age": 20,
		"phone": 595959,
	}
	map6 := make(map[string]int, 5)
	map6["age"] = 2
	map6["phone"] = 1234

	fmt.Println(map4, map5, map6) // map[age:2 phone:1234] map[age:20 phone:595959] map[age:2 phone:1234]

	//
	example := map[string]string {
		"name": "james",
		"team": "LA",
		"address": "blah",
	}

	// 순서가 없으므로 랜덤으로 값이 나옴.
	for k, v := range example {
		fmt.Println("key: ",k , "Value: ", v)
	}

	for _, v := range example {
		fmt.Println("Value: ", v)
	}

	// map 삭제
	delete(example, "address") // (변수명, key)
	fmt.Println(example) // map[name:james team:LA]

	// map 사용 시 tip

	example2 := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
	}

	value1, isExist := example2["one"] // isExist의 리턴값으로 해당키값이 map에 존재/유무를 확인할 수 있다.
	value2, isExist2 := example2["five"] // isExist의 리턴값으로 해당키값이 map에 존재/유무를 확인할 수 있다.
	_, isExist3 := example2["two"]



	fmt.Println("case1: ",value1, isExist, "case2: ", value2, isExist2) // case1: 1 true case2: 0 false

	if isExist {
		fmt.Println("exampl2에 존재!") // 실행
	} else {
		fmt.Println("없음!")
	}

	if isExist3 {
		fmt.Println("exampl2에 존재!") // 실행
	} else {
		fmt.Println("없음!")
	}

}
