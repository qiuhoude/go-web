package array

import (
	"errors"
	"fmt"
	"strings"
)

/**
 * 1)数组的插入、删除、按照下标随机访问操作
 */
type Array struct {
	data   []interface{}
	length uint
}

func New(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data: make([]interface{}, capacity, capacity),
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index > uint(cap(this.data)) {
		return true
	}
	return false
}

//通过索引查找数组，索引范围[0,n-1]
func (this *Array) Find(index uint) (interface{}, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

// 插入
func (this *Array) Insert(index uint, v interface{}) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) InsertToTail(v interface{}) error {
	return this.Insert(this.Len(), v)
}

// 删除索引index上的值
func (this *Array) Delete(index uint) (interface{}, error) {
	if index > this.Len() || this.isIndexOutOfRange(index) {
		return nil, errors.New("out of index range")
	}
	ret := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return ret, nil
}

func (this *Array) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for i := uint(0); i < this.Len(); i++ {
		sb.WriteString(fmt.Sprintf("%v", this.data[i]))
		if i != this.Len()-1 {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(']')
	return sb.String()
}
