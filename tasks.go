package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const filePath = "tasks.json"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (tasks *Tasks) get() error {
	file, err := os.Open(filePath)
	if err != nil {
		tasks = &Tasks{}
		return fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		tasks = &Tasks{}
		return fmt.Errorf("reading file: %w", err)
	}

	err = json.Unmarshal(data, tasks)
	if err != nil {
		tasks = &Tasks{}
		return fmt.Errorf("unmarshalling JSON: %w", err)
	}

	return nil
}

func (tasks *Tasks) add(description string) {
	var latestId uint
	latestId = 0
	for _, task := range tasks.Tasks {
		if task.ID > latestId {
			latestId = task.ID
		}
	}

	tasks.Tasks = append(tasks.Tasks, Task{
		ID:          latestId + 1,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	})
}

func (tasks *Tasks) update(id uint, description string) error {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Description = description
			return nil
		}
	}

	return fmt.Errorf(`task with id "%v" not found`, id)
}

func (tasks *Tasks) delete(id uint) error {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf(`task with id "%v" not found`, id)
}

func (tasks *Tasks) mark(id uint, status Status) error {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Status = status
			return nil
		}
	}

	return fmt.Errorf(`task with id "%v" not found`, id)
}

func (tasks Tasks) print() {
	if len(tasks.Tasks) > 0 {
		for _, task := range tasks.Tasks {
			fmt.Println("")
			fmt.Printf("Task ID: %d\n", task.ID)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Status: %s\n", task.Status)
			fmt.Printf("Created At: %s\n", task.CreatedAt)
			fmt.Printf("Updated At: %s\n", task.UpdatedAt)
		}
	} else {
		fmt.Println("No tasks found")
	}
}

func (tasks *Tasks) save() {
	updatedBytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, updatedBytes, 0644)
	if err != nil {
		panic(err)
	}
}
