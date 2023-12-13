package append_slice

const speculativeLength = 1

func appendSpeculativeLength(b []byte) ([]byte, int) {
	var s string
	_ = s[1:2]

	pos := len(b)
	b = append(b, "\x00\x00\x00\x00"[:speculativeLength]...)
	return b, pos
}
