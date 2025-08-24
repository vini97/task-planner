package service

import (
	"task-planner/internal/repository"
)

// TaskService provides task-related services.
type TaskService struct {
	repo *repository.Repository
}

// NewTaskService creates a new TaskService.
func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask creates a new task.
func (s *TaskService) CreateTask(name, content string) (*repository.Task, error) {
	task := &repository.Task{
		Name:    name,
		Content: content,
		Done:    false,
	}
	err := s.repo.CreateTask(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetTask retrieves a task by its ID.
func (s *TaskService) GetTask(id int64) (*repository.Task, error) {
	return s.repo.GetTask(id)
}

// GetAllTasks retrieves all tasks.
func (s *TaskService) GetAllTasks() ([]repository.Task, error) {
	return s.repo.GetTasks()
}

// UpdateTask updates a task.
func (s *TaskService) UpdateTask(id int64, name, content string, done bool) (*repository.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, err
	}

	task.Name = name
	task.Content = content
	task.Done = done

	if err := s.repo.UpdateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes a task by its ID.
func (s *TaskService) DeleteTask(id int64) error {
	return s.repo.DeleteTask(id)
}
