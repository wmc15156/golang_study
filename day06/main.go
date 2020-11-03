package main

import "fmt"

type Car struct {
	// 멤버변수 설정
	name string
	color string
	price int64
	tax int64
}

type totCost func(int, int) int

type cnt int

// 예제 4

type shoppingCart struct {
	cnt int
	price int
}

// 구조체 < - > 메소드 바인딩
// 구조체의 메소드지정할려면 함수로 만들어야함.


// 일반메소드
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체 < - > 메소드 바인딩
// func (구조체) 함수이름 리턴
func (c Car) Price() int64 {
	return c.tax + c.price
}

func describe (cnt int, price int, function totCost)  {
	fmt.Printf(" %d %d %d", cnt, price, function(cnt, price))
	fmt.Println()
}

// 예제4
// 일반 기본적인 계산
func (s shoppingCart) totPrice() int {
	return s.price * s.cnt
}

func (s *shoppingCart) totPirceP(cnt int, price int) int {
	s.cnt += cnt
	s.price += price

	return s.price * s.cnt
}

func (s shoppingCart) totPriceD(cnt int, price int) int {
	fmt.Println(s.cnt ,s.price)
	s.cnt += cnt
	s.price += price
	return s.cnt * s.price
}

func main() {
	//type 사용자 정의 타입연습

	// 사용자 정의 타입
	// GO에서는 객체지향타입을 구조체로 정의한다.(클래스 상속개념 없음)
	// 객체 지향 -> 클래스(속성: 멤버변수(색상), 기능(상태: 메소드, 전진, 후진))): 코드의 재사용성, 코드의 관리가 용이
	// GO는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스형태의 코딩가능,
	// 상태, 메소드를 분라해서 정의(결합성 없음)
	// 사용자 정의 타입: 구조체, 인터페이스, 기본타입,
	// 구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식 처럼 사용가능(객체 지향)

	bmw := Car{name: "520d", price: 5000 , color: "white", tax: 200 }
	benz := Car{name: "220e", price: 6000 , color: "white", tax: 400 }

	fmt.Println(bmw)
	fmt.Println(benz)

	fmt.Println(bmw.Price())

	// 예제2)
	// 기본자료형 사용자 정의 타입
	a := cnt(5)

	var b cnt = 15

	fmt.Println(a)
	fmt.Println(b)
	// 사용자 정의 타입 < - > 기본타입: 매개변수 전달시에 변환해야 사용가능
	// 밑에 예제 참조
	convertT(int(b))

	// convertT(b) // 에러발생 cannot use b (type cnt) as type int in argument to convertT

	// 예제3
	// 함수 사용자 정의 타입

	var orderPrice totCost

	orderPrice = func(cnt int, price int) int {
		return cnt * price * 100
	}

	describe(3, 1000, orderPrice)

	// 예제4
	// 함수는 기본적으로 값호출 => 변수의 값이 복사 후 내부 전달(원본수정(x)) -> 맵, 슬라이스 참조 전달
	// 리시버는 포인터를 활용해서 메소드내에서 원본 수정가능

	bs1 := shoppingCart{cnt: 10, price: 100}

	fmt.Println("initialize price: ", bs1.totPrice())
	fmt.Println("리시버 전달 메소드실행: ", bs1.totPirceP(10, 100)) // 4000
	fmt.Println("리시버 전달 메소드실행: ", bs1.totPirceP(10, 0)) // 6000
	fmt.Println("initialize price: ", bs1.totPrice()) // 6000
	fmt.Println("그냥메소드 실행: ", bs1.totPriceD(10, 0)) // 8000
	fmt.Println("initialize price: ", bs1.totPrice()) // 6000 원본값은 안바뀜


}

func convertT(i int) {
	fmt.Println("here Test type: ", i)
}

func convertD(i cnt) {
	fmt.Println("here Test type: ", i)
}

