// const/main.go
package main

import (
	"fmt"
)

func main() {
	// 单变量
	const a int = 1
	const b = 10

	fmt.Println("单变量：")
	fmt.Printf("变量a:%d,变量b:%d\n", a, b)

	// 块变量
	const (
		c    int64  = 5
		d    string = "Hello"
		e           = "world!"
		f, g        = "Hello ", "World!"
	)

	fmt.Println("块变量:")
	fmt.Printf("变量c:%d,变量d:%s,变量e:%s\n变量f,g:%s,%s\n", c, d, e, f, g)
}
