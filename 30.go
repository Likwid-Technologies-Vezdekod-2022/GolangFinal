package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func task30() {
	var cpu_count int
	var max_cpu_count int = runtime.NumCPU()

	fmt.Println("Введите количество процессоров\n")
	fmt.Fscan(os.Stdin, &cpu_count)

	if cpu_count > max_cpu_count {
		cpu_count = max_cpu_count
	}

	fmt.Printf("Задачи будут выполняться на %d процессорах \n", cpu_count)

	runtime.GOMAXPROCS(cpu_count)

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var taskNumber int = 0

	for scanner.Scan() {

		var activeTasksCount int = runtime.NumGoroutine()

		if len(scanner.Text()) > 0 {

			if activeTasksCount < cpu_count+1 {
				taskNumber += 1

				wg.Add(1)

				var taskDuration = scanner.Text()

				go limitTask(taskNumber, taskDuration)

			} else {

				wg.Wait()
			}

		}

	}

	wg.Wait()

	fmt.Printf("Выполнение задач закончено")

}

func limitTask(taskNumber int, taskDuration string) {
	defer wg.Done()

	taskDurationParsed, _ := time.ParseDuration(taskDuration)

	fmt.Printf("Начато выполнение задачи по номером --> %d \n", taskNumber)

	time.Sleep(taskDurationParsed)

	fmt.Printf("Закончено  выполнение задачи по номером --> %d \n", taskNumber)

}
