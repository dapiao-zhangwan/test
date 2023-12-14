package param_no_name

import (
	"sync"
	"unsafe"
)

type pointer struct{ p unsafe.Pointer }

func (p pointer) Int32() *int32 { return (*int32)(p.p) }

type coderFieldInfo struct {
}

type mergeOptions struct {
}

// 函数返回指针给指向的值赋值,赋值表达式左边无变量名
func mergeInt32(dst, src pointer, _ *coderFieldInfo, _ mergeOptions) {
	*dst.Int32() = *src.Int32()
}

func Test1() {
	var m map[int]int

	mutex := &sync.Mutex{}
	mutex.Lock()

	m[2] = 2
}

var (
	mutex1 = &sync.Mutex{}
	mutex2 = sync.Mutex{}
	mutex3 = &sync.RWMutex{}

	m1 = map[int]int{}
)

func Test3() {
	mutex1.Lock()

	Test4()

	mutex1.Unlock()
}

func Test4() {
	mutex2.Lock()

	Test1()

	mutex2.Unlock()
}

func Test5() {
	mutex2.Lock()

	Test6()

	mutex2.Unlock()
}

func Test6() {
	mutex1.Lock()

	Test1()

	mutex1.Unlock()
}

func Test7() {
	mutex3.RLock()

	mutex3.Unlock()
}

func Test8() {
	m1[2] = 3
}

func Test9() {
	_ = m1[3]
	m1[3] = 4
}

func Test10() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		Test8()
		wg.Done()
	}()

	go func() {
		Test9()
		wg.Done()
	}()

	wg.Wait()
}
