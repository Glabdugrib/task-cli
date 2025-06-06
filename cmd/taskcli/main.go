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

	validArgs, err := t.ValidateArgs(os.Args)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	t.PrintArgs(validArgs)

	taskRepository, err := t.NewRepository(filePath)
	if err != nil {
		log.Fatalf("Error initializing task repository: %v", err)
	}

	taskService := t.NewService(taskRepository)

	switch validArgs.Action {
	case t.ActionList:
		taskService.PrintTasks(validArgs.Status)
	case t.ActionAdd:
		err := taskService.CreateTask(validArgs.Description)
		if err != nil {
			log.Fatalf("Error creating task: %v", err)
		}
	case t.ActionUpdate:
		err := taskService.UpdateTaskDescription(validArgs.ID, validArgs.Description)
		if err != nil {
			log.Fatalf("Error updating task: %v", err)
		}
	case t.ActionMark:
		err := taskService.UpdateTaskStatus(validArgs.ID, *validArgs.Status)
		if err != nil {
			log.Fatalf("Error marking task: %v", err)
		}
	case t.ActionDelete:
		err := taskService.DeleteTask(validArgs.ID)
		if err != nil {
			log.Fatalf("Error deleting task: %v", err)
		}
	}
}
