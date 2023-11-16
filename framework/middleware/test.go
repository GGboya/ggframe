package middleware

import (
	"fmt"
	"ggframe/framework"
)

func Test1() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next() // 调用next往下调用，自增ctx.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next() // 调用next往下调用，自增ctx.index
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		c.Next() // 调用next往下调用，自增ctx.index
		fmt.Println("middleware post test3")
		return nil
	}
}
