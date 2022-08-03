package main

import (
	"fmt"

	"golang.org/x/example/stringutil"

	"me/mymodule"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))
	fmt.Println(mymodule.Fun())
}
