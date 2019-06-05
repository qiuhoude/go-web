package demo2_mysql

import "testing"

func TestOpen(t *testing.T) {
	open()
}

func TestInsert(t *testing.T) {
	insert()
}

func TestUpdate(t *testing.T) {
	update()
}

func TestQueryOne(t *testing.T) {
	queryOne()
}
func TestQueryMulti(t *testing.T) {
	queryMulti()
}

func TestTransaction(t *testing.T) {
	transaction()
}
