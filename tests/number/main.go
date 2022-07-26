// +build gofuzz
package fuzz

import "github.com/ezoic/parse"

func Fuzz(data []byte) int {
	_ = parse.Number(data)
	return 1
}
