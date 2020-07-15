package shopee

import "fmt"

func f1(a []int) {
	a[0] = 100
	a = append(a, 30)
}

func TestDiDi() {
	a := make([]int, 2)
	a[0] = 10
	a[1] = 20

	f1(a)
	fmt.Println(a)
}
