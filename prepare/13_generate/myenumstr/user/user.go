package user

type Status int

//go:generate go run ../myenumstr.go -type Status,Color

const (
	Offline Status = iota
	Online
	Disable
	Deleted
)

type Color int

const (
	Write Color = iota
	Red
	Blue
)
