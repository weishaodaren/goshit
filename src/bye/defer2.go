package main

import "fmt"

func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to db", 1)
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db", 6)
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.", 2)
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...", 3)
	fmt.Println("Oops! some crash or network error ...", 4)
	fmt.Println("Returning from function here!", 5)
	return
}
