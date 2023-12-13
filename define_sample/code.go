package define_sample

import (
	"math/rand"
	"unsafe"
)

var v1 int
var v2 string

var int1, int2 int

var (
	str1 string
	b1   bool
)

var ir1 = rand.Int()

var ir2, ir3 = rand.Int(), rand.Int()

var sc1, sc2 = create1()

func create1() (string, string) {
	return "", ""
}

var scp1 = sc1

type s1 struct {
}

type s2 s1

func (p *s1) f1() {

}

type m1 map[string]int

func (p m1) f1() {

}

type i1 interface {
	f1()
}

type func1 func(i int, j int) (str string)

type inttype1 int

type chantype1 chan string

type interface1 interface{}

type arrtype1 []int

type i11 i1

type func11 func1

type p1 *int

func test(v1 int) {
	for v1 := 0; v1 < 100; v1++ {
		for i := 0; i < 20; i++ {
			var v1 string
			v2 = v1
		}

		var v1 float32
		_ = v1
	}
}

func test2() {
	test(2)
	rand.Intn(1)
	m11 := m1{}
	m11.f1()
	sc2 = "s"
}

func test3(m map[string]map[int]map[string]int, s [][]map[int]int, f func11, c1 chan chan int, c2 chan arrtype1, p unsafe.Pointer, it1 interface1, i11 i1, s1 ...int) interface{} {
	return 2
}
