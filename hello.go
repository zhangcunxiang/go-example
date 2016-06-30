package main

import (
	"fmt"
	"github.com/user/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("Hello, world.\n"))
	fmt.Printf(stringutil.Reverse("n你好，哈哈啊哈啊，😄"))
	i := 43
	fmt.Printf(" i is the type of %T", i)
	f := float64(i)
	fmt.Printf(" f is the type of %T", f)
}
