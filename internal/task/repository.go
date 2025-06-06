package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

type TaskRepository struct {
	tasks     []Task
	storePath string
}

// Creates a new task repository that persists to a JSON file
func NewRepository(path string) (*TaskRepository, error) {
	repo := &TaskRepository{storePath: path}
	if err := repo.load(); err != nil {
		return nil, err
	}
	return repo, nil
}

// List retrieves all tasks, optionally filtering by status
func (r *TaskRepository) List(status *Status) []Task {
	if status == nil {
		return r.tasks
	}

	var filtered []Task
	for _, task := range r.tasks {
		if task.Status == *status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// Create a new task with a unique ID
func (r *TaskRepository) Add(task Task) error {
	task.ID = r.nextID()
	r.tasks = append(r.tasks, task)
	return r.save()
}

// Modifies an existing task by ID
func (r *TaskRepository) Update(updatedTask Task) error {
	for i, t := range r.tasks {
		if t.ID == updatedTask.ID {
			r.tasks[i] = updatedTask
			return r.save()
		}
	}
	return fmt.Errorf("task not found")
}

// Deletes a task by ID
func (r *TaskRepository) Delete(id uint) error {
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = slices.Delete(r.tasks, i, i+1)
			return r.save()
		}
	}
	return fmt.Errorf("task not found")
}

// Generates a new unique ID for a task
func (r *TaskRepository) nextID() uint {
	maxID := uint(0)
	for _, t := range r.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// Saves the current state of tasks to the JSON file
func (r *TaskRepository) save() error {
	data, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.storePath, data, 0644)
}

// Loads tasks from the JSON file, creating it if it doesn't exist
func (r *TaskRepository) load() error {
	if _, err := os.Stat(r.storePath); errors.Is(err, os.ErrNotExist) {
		// Create empty file if it doesn't exist
		return os.WriteFile(r.storePath, []byte("[]"), 0644)
	}
	data, err := os.ReadFile(r.storePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &r.tasks)
}
