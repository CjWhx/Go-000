package main

import (
	errs "errors"
	"fmt"
	pkg_errs "github.com/pkg/errors"
)

func main() {
	err := errs.New("this is error")
	fmt.Printf("%T %v \n", err, err)

	fmt.Println("----------")

	// Wrap 会带有调用栈信息
	err1 := pkg_errs.Wrap(err, "this is wrapped!")
	pkg_err := pkg_errs.WithMessage(err1, "this is with")
	fmt.Printf("print err, type: %T value:%+v", pkg_err, pkg_err)
}
