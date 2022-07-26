// +build gofuzz
package fuzz

import "github.com/ezoic/parse"

func Fuzz(data []byte) int {
	_, _ = parse.Mediatype(data)
	return 1
}
