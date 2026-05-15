package main

import "sync"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	mu    sync.Mutex
	tasks []Task
)

func AddTask(task Task) {
	mu.Lock()
	defer mu.Unlock()

	tasks = append(tasks, task)
}

func GetTasks() []Task {
	return tasks
}
