package main

import "fmt"

func main() {
	i := 3
	fmt.Printf("当 i = %d 时 \n", i)

	// 默认携带 break
	switch i {
	case 1:
		fmt.Println("输出 i= ", 1)
	case 2:
		fmt.Println("输出 i= ", 2)
	case 3:
		fmt.Println("输出 i= ", 3)
		fallthrough // 不跳出 执行下一个
	case 4, 5, 6: // 可以有多个选项
		fmt.Println("输出 i= ", 4, 5, 6)
	default:
		fmt.Println("xxx")
	}

}
