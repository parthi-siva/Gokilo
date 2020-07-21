package main

import "fmt"

const (
	HL_NORMAL int = iota
	HL_NONPRINT
	HL_COMMENT
	HL_MLCOMMENT
	HL_KEYWORD1
	HL_KEYWORD2
	HL_STRING
	HL_NUMBER
	HL_MATCH
)

const (
	HL_HIGHLIGHT_STRING = 1 << 0
	HL_HIGHLIGHT_NUMBER = 1 << 1
)

func main() {
	fmt.Println()
}
