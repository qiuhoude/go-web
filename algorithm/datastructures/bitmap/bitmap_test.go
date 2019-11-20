package bitmap

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	bm := NewBitMap(int(1 * KB))
	bm.Set(100)
	bm.Set(20)
	bm.Set(10)
	bm.Set(9001)

	assert.Equal(t, bm.Get(100), true)
	assert.Equal(t, bm.Get(20), true)
	assert.Equal(t, bm.Get(10), true)
	assert.Equal(t, bm.Get(21), false)
	//assert.Equal(t, bm.Get(9001), false)

}
