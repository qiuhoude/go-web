package demo3_redis

import "testing"

func TestOpen(t *testing.T) {
	open()
}

func TestGetSet(t *testing.T) {
	getSet()
}

func TestMgetmset(t *testing.T) {
	mgetmset()
}
func TestList(t *testing.T) {
	list()
}
func TestHash(t *testing.T) {
	hash()
}

func TestPipelining(t *testing.T) {
	pipelining()
}
