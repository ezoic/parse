// +build gofuzz
package fuzz

import "github.com/ezoic/parse"

func Fuzz(data []byte) int {
	_, _, _ = parse.DataURI(data)
	return 1
}
