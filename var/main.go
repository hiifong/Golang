// var/main.go
package main

import "fmt"

func main() {
	// 单变量
	var a int64 // 零值：0
	var b string = "分"
	var c = "100分"

	// 块变量
	var (
		d    string // 零值：""
		e    string = "e is string type"
		f           = "this is string"
		g, h        = 100, "分"
	)
	fmt.Println("单变量：")
	fmt.Printf("变量a:%d,变量b:%s,变量c:%s\n", a, b, c)
	fmt.Println("块变量：")
	fmt.Printf("变量d:%s,变量e:%s,变量f:%s,变量g:%d,变量h:%s\n", d, e, f, g, h)
}
