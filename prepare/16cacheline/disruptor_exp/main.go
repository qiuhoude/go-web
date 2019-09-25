package main

import (
	"fmt"
	"github.com/smartystreets/go-disruptor"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	BufferSize   = 1024 * 64
	BufferMask   = BufferSize - 1
	Iterations   = 10
	Reservations = 16
)

var (
	ring = [BufferSize]int64{}
)

type SampleConsumer struct {
	id int
}

func (s SampleConsumer) Consume(lower, upper int64) {
	gid := GoID()
	for ; lower <= upper; lower++ {
		message := ring[lower&BufferMask]
		fmt.Printf("goid:%d consumeId:%d msg:%d \n", gid, s.id, message)
	}
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	controller := disruptor.Configure(BufferSize).WithConsumerGroup(SampleConsumer{1}, SampleConsumer{2}).Build()

	controller.Start()

	started := time.Now()
	publish(controller.Writer())
	finished := time.Now()
	select {}
	//controller.Stop()
	fmt.Println(Iterations, finished.Sub(started))
}

func publish(writer *disruptor.Writer) {
	sequence := disruptor.InitialSequenceValue
	for i := 0; i < Iterations; i++ {
		sequence = writer.Reserve(1) // java api中的next
		ring[sequence&BufferMask] = int64(i)
		writer.Commit(sequence, sequence)
	}

}

// func publish(writer *disruptor.Writer) {
// 	sequence := disruptor.InitialSequenceValue
// 	for sequence <= Iterations {
// 		sequence += Reservations // only an advantage at smaller reservations, e.g. 1-4?
// 		writer.Await(sequence)
// 		for lower := sequence - Reservations + 1; lower <= sequence; lower++ {
// 			ring[lower&BufferMask] = lower
// 		}
// 		writer.Commit(sequence-Reservations+1, sequence)
// 	}
// }
