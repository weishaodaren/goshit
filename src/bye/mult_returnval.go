package main

func main() {
	num1, num2, num3 := SetReturnVal(10, 5)
	println(num1, num2, num3)

	num1, num2, num3 = SetReturnVal_2(10, 5)
	println(num1, num2, num3)
}

func SetReturnVal(input1, input2 int) (int, int, int) {
	return input1 + input2, input1 * input2, input1 - input2
}

func SetReturnVal_2(input1, input2 int) (s1, s2, s3 int) {
	s1, s2, s3 = input1 + input2, input1 * input2,  input1 - input2
	return
}
