package main

import (
	"fmt"
	"github.com/guosaitong/foundation/response"
)

func main() {
	// 成功案例
	result1 := response.GetUser("123")
	fmt.Printf("成功返回: %+v\n", result1)

	// 错误案例1
	result2 := response.GetUser("")
	fmt.Printf("参数错误: %+v\n", result2)

	// 错误案例2
	result3 := response.GetUser("456")
	fmt.Printf("查询失败: %+v\n", result3)
}
