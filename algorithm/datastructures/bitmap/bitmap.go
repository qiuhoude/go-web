package bitmap

type BitMap struct {
	data []byte
	len  int
}

// n表示存储数据的大小
func NewBitMap(n int) *BitMap {
	return &BitMap{
		data: make([]byte, n/8+1),
		len:  n,
	}
}

const bitSize = 8

type unit int

const (
	Bit unit = 1
	B        = 8 * Bit
	KB       = 1024 * B
	MB       = 1024 * KB
	GB       = 1024 * MB
	TB       = 1024 * GB
)

func (b *BitMap) Set(k uint) {
	if k > uint(b.len) { //超过存储范围
		return
	}

	index := k / bitSize
	var bit uint
	bit = uint(k) % bitSize
	b.data[index] |= 1 << bit
}

func (b *BitMap) Get(k uint) bool {
	if k > uint(b.len) { //超过存储范围
		return false
	}
	index := k / bitSize
	var bit uint
	bit = uint(k) % bitSize
	return b.data[index]&(1<<bit) != 0
}
