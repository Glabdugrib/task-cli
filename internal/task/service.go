package task

import (
	"fmt"
	"time"
)

type TaskService struct {
	repo *TaskRepository
}

func NewService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() []Task {
	return s.repo.List()
}

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

	return s.repo.Add(task)
}

func (s *TaskService) UpdateTaskDescription(id uint, description string) error {
	tasks := s.repo.List()
	for _, t := range tasks {
		if t.ID == id {
			t.Description = description
			return s.repo.Update(t)
		}
	}
	return fmt.Errorf("task not found")
}

func (s *TaskService) UpdateTaskStatus(id uint, status Status) error {
	tasks := s.repo.List()
	for _, t := range tasks {
		if t.ID == id {
			t.Status = status
			return s.repo.Update(t)
		}
	}
	return fmt.Errorf("task not found")
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

func (s *TaskService) PrintTasks() {
	tasks := s.repo.List()

	for _, t := range tasks {
		fmt.Printf("\nTASK %v\n", t.ID)
		fmt.Printf("Description: %s\n", t.Description)
		fmt.Printf("Status: %s\n", t.Status.String())
		fmt.Printf("Created at: %s\n", t.CreatedAt)
		fmt.Printf("Updated at: %s\n", t.UpdatedAt)
	}
}
