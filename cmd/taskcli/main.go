package main

import (
	"fmt"
	t "go-task-list/internal/task"
	"log"
	"os"
)

const filePath = "tasks.json"

func main() {
	fmt.Println("Welcome to Golang Task CLI!")

	parsed, err := t.ValidateArgs(os.Args)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	taskRepository, err := t.NewRepository(filePath)
	if err != nil {
		log.Fatalf("Error initializing task repository: %v", err)
	}

	taskService := t.NewService(taskRepository)

	switch parsed.Action {
	case t.ActionList:
		taskService.PrintTasks()
	case t.ActionAdd:
		err := taskService.CreateTask(parsed.Description)
		if err != nil {
			log.Fatalf("Error creating task: %v", err)
		}
	case t.ActionUpdate:
		err := taskService.UpdateTaskDescription(parsed.ID, parsed.Description)
		if err != nil {
			log.Fatalf("Error updating task: %v", err)
		}
	case t.ActionMark:
		err := taskService.UpdateTaskStatus(parsed.ID, parsed.Status)
		if err != nil {
			log.Fatalf("Error marking task: %v", err)
		}
	case t.ActionDelete:
		err := taskService.DeleteTask(parsed.ID)
		if err != nil {
			log.Fatalf("Error deleting task: %v", err)
		}
	}
}
