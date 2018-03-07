package iotas

const (
	_  = iota //iota 0 assign it to blank _ identifier
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)
