package main

import (
	"fmt"
	"github.com/juancho85/02_packages/stringutils"
)

func main() {
	fmt.Println(stringutils.MyName)
	fmt.Println(stringutils.Reverse(stringutils.MyName))
}
