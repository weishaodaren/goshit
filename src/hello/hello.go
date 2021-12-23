package main

import ("fmt")

func main(){
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2m string = "Tom", "Jay"
	fmt.Println(name_1, name_2m)

	const name_3, _age_1 = "TOM", 30
	fmt.Println(name_3, _age_1)

	var age__1 uint8 = 32
	age_3 := 23

	fmt.Println(age__1, age_3)
}