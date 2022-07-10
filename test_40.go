package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

var allTasksDuration int = 0

func test_40() {

	test_range := 30

	for i := 1; i < test_range; i++ {
		wg.Add(4)

		allTasksDuration += i
		var duration string = strconv.Itoa(i) + "s"
		go testTaskRunning(duration, i)
		go testAsyncPost(duration)
		go testSchedule()
		go testTimer()

	}

	wg.Wait()

}

func testAsyncPost(duration string) {

	defer wg.Done()

	message := map[string]interface{}{
		"duration": duration,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://127.0.0.1:8000/add?async=1", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

}

func testTaskRunning(duration string, id int) {

	defer wg.Done()

	taskDurationParsed, _ := time.ParseDuration(duration)

	time.Sleep(taskDurationParsed)

	allTasksDuration -= id

}

func testSchedule() {
	defer wg.Done()

	resp, err := http.Get("http://127.0.0.1:8000/schedule")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseSchedule, _ := strconv.Atoi(string(body))

	if responseSchedule == runtime.NumGoroutine() {
		println("Тест расписания пройден")
	}

}

func testTimer() {
	defer wg.Done()

	resp, err := http.Get("http://127.0.0.1:8000/time")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseTime, _ := strconv.Atoi(string(body))

	if responseTime == allTasksDuration {
		println("Тест таймера пройден")
	}
}
