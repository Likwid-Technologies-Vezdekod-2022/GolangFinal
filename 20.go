package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func task20() {

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

			taskNumber += 1

			wg.Add(1)

			var taskDuration = scanner.Text()

			go paralelRunTask(taskNumber, taskDuration)

		}

	}

	wg.Wait()

	fmt.Printf("Выполнение задач закончено")

}

func paralelRunTask(taskNumber int, taskDuration string) {

	defer wg.Done()

	taskDurationParsed, _ := time.ParseDuration(taskDuration)

	fmt.Printf("Начато выполнение задачи по номером --> %d \n", taskNumber)

	time.Sleep(taskDurationParsed)

	fmt.Printf("Закончено  выполнение задачи по номером --> %d \n", taskNumber)

}
