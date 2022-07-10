package main

import "fmt"

func main() {
	fmt.Println("---------------------------------------\n")
	fmt.Println("Задание за 10 баллов ")
	task10()
	fmt.Println("---------------------------------------\n")

	fmt.Println("Задание за 20 баллов ")
	task20()
	fmt.Println("---------------------------------------\n")

	fmt.Println("Задание за 30 баллов ")
	task30()
	fmt.Println("---------------------------------------\n")

	go task40()
	fmt.Println("Задание за 40 баллов ")
	test_40()
	fmt.Println("---------------------------------------\n")

	wg.Wait()
}
