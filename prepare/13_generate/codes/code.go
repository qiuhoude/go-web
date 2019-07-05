package codes

type Code uint32

//go:generate stringer -type Code
// go:generate java -version

const (
	OK Code = iota
	Canceled
	Unknown
	InvalidArgument
)
