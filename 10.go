package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var taskNumber int = 0

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fmt.Printf("---------------------------------------- \n")

			taskNumber += 1

			fmt.Printf("Начато выполнение задачи по номером --> %d \n", taskNumber)

			var taskDuration = scanner.Text()

			task(taskDuration)

			fmt.Printf("Закончено  выполнение задачи по номером --> %d \n", taskNumber)

			fmt.Printf("---------------------------------------- \n")
		}

	}

	fmt.Println("Выполнение задач закончено")

}

func task(taskDuration string) {

	taskLengthParsed, _ := time.ParseDuration(taskDuration)

	time.Sleep(taskLengthParsed)

}
