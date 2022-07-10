package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID       string `json:"id"`
	Duration string `json:"duration"`
}

var tasks []Task

func task40() {
	r := mux.NewRouter()

	r.HandleFunc("/time", getQueueTime).Methods("GET")
	r.HandleFunc("/add", addTask).Methods("POST")
	r.HandleFunc("/schedule", getSchedule).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getQueueTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(tasks) <= 0 {
		w.WriteHeader(400)
		w.Write([]byte("В очереди нет доступных задач"))
		return
	}

	allQueueDuration, _ := time.ParseDuration(tasks[0].Duration)

	for _, task := range tasks {
		taskDuration, _ := time.ParseDuration(task.Duration)
		allQueueDuration += taskDuration
	}

	json.NewEncoder(w).Encode(allQueueDuration.Seconds())
}

func getSchedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inQueueTasksCount int = len(tasks)

	json.NewEncoder(w).Encode(inQueueTasksCount)
}

func addTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sync, sync_ok := r.URL.Query()["sync"]

	async, async_ok := r.URL.Query()["async"]

	if !async_ok && !sync_ok {
		w.WriteHeader(400)
		w.Write([]byte("Не передан параметр async или sync"))
		return
	}

	if async_ok && sync_ok {
		json.NewEncoder(w).Encode("Переданы одновременно два параметра, попробуйте еще раз (передавать можно только 1)")
		return
	}

	var task Task

	task.ID = strconv.Itoa(rand.Intn(1000000))

	_ = json.NewDecoder(r.Body).Decode(&task)

	tasks = append(tasks, task)

	json.NewEncoder(w).Encode(task)

	if sync_ok {

		if sync[0] == "1" {

			wg.Add(1)
			go startTask(task)

			wg.Wait()
		}
	}

	if async_ok {
		if async[0] == "1" {
			wg.Add(1)
			go startTask(task)

		}
	}

	return
}

func startTask(task Task) {

	defer wg.Done()

	taskDurationParsed, _ := time.ParseDuration(task.Duration)

	fmt.Printf("Начато выполнение задачи с id --> %s \n", task.ID)

	time.Sleep(taskDurationParsed)

	fmt.Printf("Закончено  выполнение задачи с id --> %s \n", task.ID)

	deleteTaskFromQueue(task.ID)

}

func deleteTaskFromQueue(taskId string) {
	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

}
