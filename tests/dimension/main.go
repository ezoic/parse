// +build gofuzz
package fuzz

import "github.com/ezoic/parse"

func Fuzz(data []byte) int {
	_, _ = parse.Dimension(data)
	return 1
}
