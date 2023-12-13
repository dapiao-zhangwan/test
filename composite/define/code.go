package define

import "sync"

type DefineS struct {
	I1 int
	S1 string
}

func Test1() {
	var m map[int]int

	mutex := &sync.Mutex{}
	mutex.Lock()

	m[2] = 2
}
