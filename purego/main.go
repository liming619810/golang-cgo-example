package main

import (
	"fmt"

	"github.com/ebitengine/purego"
)

func main() {
	libc, err := openLibrary("test_so.so")
	if err != nil {
		panic(err)
	}

	var test_so_func func(a int, b int) int
	purego.RegisterLibFunc(&test_so_func, libc, "test_so_func") //注册函数
	fmt.Println(test_so_func(2, 2))                             //执行函数
}

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
