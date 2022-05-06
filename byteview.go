package go_cache

type ByteView struct {
	b []byte
}

func (b ByteView) Len() int {
	return len(b.b)
}

func (b *ByteView) ByteSlice() []byte {
	return cloneBytes(b.b)
}

func (b *ByteView) String() string {
	return string(b.b)
}

// b是只读的，返回一个copy，防止缓存值被外部修改
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
