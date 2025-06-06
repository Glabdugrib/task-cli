package task

import (
	"fmt"
	"time"
)

type TaskService struct {
	repo *TaskRepository
}

// Creates a new task service with the given repository
func NewService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// Retrieves all tasks
func (s *TaskService) GetTasks() []Task {
	return s.repo.List(nil)
}

// Creates a new task with the given description
func (s *TaskService) CreateTask(description string) error {
	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	currentTime := time.Now().Format(time.RFC3339)
	task := Task{
		Description: description,
		Status:      StatusPending,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	err := s.repo.Add(task)
	if err != nil {
		return err
	}

	// Only reached if no error
	fmt.Println("\nTask created successfully")
	return nil
}

// Updates the description of an existing task by ID
func (s *TaskService) UpdateTaskDescription(id uint, description string) error {
	tasks := s.repo.List(nil)
	for _, t := range tasks {
		if t.ID == id {
			t.Description = description
			err := s.repo.Update(t)
			if err != nil {
				return err
			}

			fmt.Println("\nTask updated successfully")
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

// Updates the status of an existing task by ID
func (s *TaskService) UpdateTaskStatus(id uint, status Status) error {
	tasks := s.repo.List(nil)
	for _, t := range tasks {
		if t.ID == id {
			t.Status = status
			err := s.repo.Update(t)
			if err != nil {
				return err
			}

			fmt.Println("\nTask updated successfully")
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

// Deletes a task by ID
func (s *TaskService) DeleteTask(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	fmt.Println("\nTask deleted successfully")
	return nil
}

// Prints the details of all tasks, optionally filtered by status
func (s *TaskService) PrintTasks(status *Status) {
	fmt.Println("\nTASKS LIST:")
	tasks := s.repo.List(status)

	if len(tasks) == 0 {
		fmt.Println("\nNo tasks found.")
		return
	}

	for _, t := range tasks {
		fmt.Printf("\nID %v\n", t.ID)
		fmt.Printf("Description: %s\n", t.Description)
		fmt.Printf("Status: %s\n", t.Status.String())
		fmt.Printf("Created at: %s\n", t.CreatedAt)
		fmt.Printf("Updated at: %s\n", t.UpdatedAt)
	}
}
