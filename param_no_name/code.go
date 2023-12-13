package param_no_name

import "unsafe"

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
