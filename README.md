# Task Planner API

A simple yet powerful task manager to help you organize your day and focus on what's important.

## Features

- Create, read, update, and delete tasks (CRUD).
- RESTful API interface for easy integration.
- Data persistence with PostgreSQL.

## Getting Started

Follow these instructions to get a copy of the project running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.20 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/task-planner.git
    cd task-planner
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Set up environment variables:**

    Create a `.env` file in the project root and add your PostgreSQL database password:

    ```env
    DB_PASSWORD="your-secret-password"
    ```

4.  **Run the application:**
    ```bash
    go run cmd/tasks/main.go
    ```

    The server will start on `http://localhost:8000`.

## API Endpoints

The API provides the following endpoints to manage tasks:

| Method   | Endpoint         | Description                             |
| :------- | :--------------- | :-------------------------------------- |
| `GET`    | `/v1/tasks`      | Returns a list of all tasks.            |
| `POST`   | `/v1/tasks`      | Creates a new task.                     |
| `GET`    | `/v1/tasks/{id}` | Returns a specific task by its ID.      |
| `PUT`    | `/v1/tasks/{id}` | Updates an existing task.               |
| `DELETE` | `/v1/tasks/{id}` | Deletes a task.                         |

### Usage Examples with `curl`

**Create a task:**

```bash
curl --request POST \
  --url http://localhost:8000/v1/tasks \
  --header 'Content-Type: application/json' \
  --data 
    "{
    "name": "My First Task",
    "content": "Sample task content."
}"
```

**List all tasks:**

```bash
curl http://localhost:8000/v1/tasks
```

**Get a task by ID:**

```bash
curl http://localhost:8000/v1/tasks/1
```

**Update a task:**

```bash
curl --request PUT \
  --url http://localhost:8000/v1/tasks/1 \
  --header 'Content-Type: application/json' \
  --data 
    "{
    "name": "Updated Task",
    "content": "This task has been updated.",
    "done": true
}"
```

**Delete a task:**

```bash
curl --request DELETE http://localhost:8000/v1/tasks/1
```

## Technologies Used

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [github.com/lib/pq](https://github.com/lib/pq) - PostgreSQL driver for Go.
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - For loading environment variables from a `.env` file.
