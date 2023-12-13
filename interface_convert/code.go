package interface_convert

import "unsafe"

func dereferenceAppend(result []interface{}, arr []unsafe.Pointer, from, to uint32) []interface{} {
	for i := from; i < to; i++ {
		// We have mutual exclusion on arr, there's no need for atomics.
		x := (*interface{})(arr[i])
		if x != nil {
			result = append(result, *x)
		}

		_ = (interface{})(i)
		_ = interface{}(i)
	}
	return result
}
