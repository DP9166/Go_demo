package main

import (
	"calc/simplemath"
	"fmt"
	"os"
	"strconv"
)

// 定义一个用户打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments]...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

func main()  {
	// 用于获取命令行参数
	args := os.Args
	// 除程序名以外，至少需要传两个参数
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	// 第二个参数表示计算方法
	switch args[1] {
	case "add":
		// 至少需要包含四个参数
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1></integer2>")
			return
		}
		// 获取待相加的数值， 并将类型转化为整形
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		// 获取参数出错, 则退出
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1></integer2>")
			return
		}
		// 从simplemath 包中引入Add 方法进行加法计算
		ret := simplemath.Add(v1, v2)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	case "sqrt":
		// 至少需要包含3个参数
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 获取待计算平方根的数值, 并将类型转化为整数
		v, err := strconv.Atoi(args[2])
		// 获取参数出错, 则退出
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		// 打印计算结果
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}

}