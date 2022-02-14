package main

// 应当只使用正序的标签（标签位于 goto 语句之后）注意标签和goto之间不能出现定义变量的语句
func main() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
