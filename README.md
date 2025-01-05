# Gin Todo API

This is a simple Todo API built with Gin and PostgreSQL.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/gin-todo-api.git
    cd gin-todo-api
    ```

2. Copy the sample environment file and update it with your own values:
    ```sh
    cp .env.sample .env
    ```

3. Start the application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

4. The API will be available at [http://localhost:8080](http://_vscodecontentref_/0).

## Endpoints

- `POST /signup` - Create a new account
- `POST /login` - Login and get a JWT token
- `GET /accounts` - Get account details (requires authentication)
- `PATCH /accounts` - Update account details (requires authentication)
- `DELETE /accounts` - Delete account (requires authentication)
- `GET /tasks` - Get all tasks (requires authentication)
- `GET /tasks/:id` - Get a specific task (requires authentication)
- `POST /tasks` - Create a new task (requires authentication)
- `PATCH /tasks/:id` - Update a task (requires authentication)
- `DELETE /tasks/:id` - Delete a task (requires authentication)
