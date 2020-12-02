package main

import (
	"fmt"
	"geek.com/lesson3-4/homework/model"
	"geek.com/lesson3-4/homework/service"
)

func main() {
	userPage()
}

// 用户页面 获取所有的用户
func userPage() {
	defer func() {
		pErr := recover()
		if pErr != nil {
			fmt.Printf("程序异常退出，%+v", pErr)
			return
		}
		fmt.Println("程序正常退出")
	}()

	fmt.Print("开始查询数据...")

	persons := make([]model.User, 0)
	persons, err := service.GetAllUser()
	if err != nil {
		fmt.Printf("获取数据数据时失败,stack trace:%+v", err)
	}
	fmt.Println("查询到的数据为:", persons)
}
