package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(utf8Index("我go这是中午啊", "中午"))
	fmt.Println(strings.Index("go这是中午啊", "中午"))
}

func utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
