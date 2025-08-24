# Task Planner API

Um simples, mas poderoso, gerenciador de tarefas para ajudar a organizar seu dia e focar no que é importante.

## Features

- Crie, leia, atualize e delete tarefas (CRUD).
- Interface de API RESTful para fácil integração.
- Persistência de dados com PostgreSQL.

## Primeiros Passos

Siga estas instruções para obter uma cópia do projeto em execução em sua máquina local para desenvolvimento e teste.

### Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.20 ou superior)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/downloads)

### Instalação

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/seu-usuario/task-planner.git
    cd task-planner
    ```

2.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Configure as variáveis de ambiente:**

    Crie um arquivo `.env` na raiz do projeto e adicione a senha do seu banco de dados PostgreSQL:

    ```env
    DB_PASSWORD="sua-senha-secreta"
    ```

4.  **Execute a aplicação:**
    ```bash
    go run cmd/tasks/main.go
    ```

    O servidor será iniciado em `http://localhost:8000`.

## Endpoints da API

A API fornece os seguintes endpoints para gerenciar tarefas:

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `GET` | `/v1/tasks` | Retorna uma lista de todas as tarefas. |
| `POST` | `/v1/tasks` | Cria uma nova tarefa. |
| `GET` | `/v1/tasks/{id}` | Retorna uma tarefa específica pelo seu ID. |
| `PUT` | `/v1/tasks/{id}` | Atualiza uma tarefa existente. |
| `DELETE` | `/v1/tasks/{id}` | Exclui uma tarefa. |

### Exemplos de Uso com `curl`

**Criar uma tarefa:**

```bash
curl --request POST \
  --url http://localhost:8000/v1/tasks \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Minha Primeira Tarefa",
    "content": "Conteúdo da tarefa de exemplo."
}'
```

**Listar todas as tarefas:**

```bash
curl http://localhost:8000/v1/tasks
```

**Obter uma tarefa por ID:**

```bash
curl http://localhost:8000/v1/tasks/1
```

**Atualizar uma tarefa:**

```bash
curl --request PUT \
  --url http://localhost:8000/v1/tasks/1 \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Tarefa Atualizada",
    "content": "Esta tarefa foi atualizada.",
    "done": true
}'
```

**Excluir uma tarefa:**

```bash
curl --request DELETE http://localhost:8000/v1/tasks/1
```

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [github.com/lib/pq](https://github.com/lib/pq) - Driver do PostgreSQL para Go.
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Para carregar variáveis de ambiente de um arquivo `.env`.