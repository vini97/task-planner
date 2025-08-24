package repository

import (
	"database/sql"
	"errors"
)

// Task represents a task in the system.
type Task struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

// ErrNotFound is returned when a task is not found.
var ErrNotFound = errors.New("task not found")

// Repository handles the database operations for tasks.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new task repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// CreateTable creates the tasks table if it doesn't exist.
func (r *Repository) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		content TEXT,
		done BOOLEAN NOT NULL DEFAULT false
	);
	`
	_, err := r.db.Exec(query)
	return err
}

// CreateTask adds a new task to the database.
func (r *Repository) CreateTask(task *Task) error {
	query := "INSERT INTO tasks (name, content, done) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, task.Name, task.Content, task.Done).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetTask retrieves a task by its ID.
func (r *Repository) GetTask(id int64) (*Task, error) {
	task := &Task{}
	query := "SELECT id, name, content, done FROM tasks WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Name, &task.Content, &task.Done)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return task, nil
}

// GetTasks retrieves all tasks from the database.
func (r *Repository) GetTasks() ([]Task, error) {
	query := "SELECT id, name, content, done FROM tasks"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Content, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTask updates an existing task in the database.
func (r *Repository) UpdateTask(task *Task) error {
	query := "UPDATE tasks SET name = $1, content = $2, done = $3 WHERE id = $4"
	_, err := r.db.Exec(query, task.Name, task.Content, task.Done, task.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask removes a task from the database.
func (r *Repository) DeleteTask(id int64) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}