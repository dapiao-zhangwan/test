package composite

import "gitlab.gg.com/game_framework/code_check/samples/single_test/composite/define"

func Test1() {
	m := map[string][]int{"array": {1, 2}}["ss"]
	_ = m
}

func Test2() {
	arr := []map[string]interface{}{{
		"foo": "bar",
	}, {
		"bar": "foo",
	}}[1]

	_ = arr
}

var (
	i1 int
	s1 string

	I1 int
)

type S1 struct {
	i1 int
	s1 string
}

func Test3() {
	s := &S1{i1: 2, s1: ""}
	_ = s
}

func Test4() {
	s := &define.DefineS{I1: 2, S1: ""}
	_ = s
}
