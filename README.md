# 📝 Task API with Go, Gin & GORM

A simple and elegant RESTful API for managing tasks — written in Go with the Gin framework, backed by SQLite, and documented using Swagger.

---

## 🚀 Features

- ✅ Create, Read, Update, Delete tasks
- 🔁 Status toggle (incomplete/complete)

---

## 🧱 Tech Stack

- [Go 1.18+](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/swag) (Swagger Generator)
- SQLite (in-memory)
- Docker
- Testify (for unit testing)

---

## 📬 API Endpoints

| Method | Endpoint        | Description        |
|--------|-----------------|--------------------|
| GET    | `/tasks`        | Get all tasks      |
| POST   | `/tasks`        | Create new task    |
| PUT    | `/tasks/{id}`   | Update a task      |
| DELETE | `/tasks/{id}`   | Delete a task      |

📖 Full API documentation available at (remember to run the server first) :  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 📦 Local Development

### 🔧 Prerequisites

- Go 1.18+
- (Optional) Docker

### ✅ 1. Clone the repo

```bash
git clone https://github.com/barneeyyu/task-api.git
cd task-api
```

### 🔑 2. Install the dependencies

```bash
go mod tidy
```

### 🔑 3. Run the server

```bash
go run main.go
```

### 🔑 4. Run the tests

```bash
go test
```

## 🐳 Docker Deployment

If you prefer to run the app in a container:

### 🛠️ Build Docker image

```bash
docker build -t task-api .
```

### 🚀 Run Docker container

```bash
docker run -d -p 8080:8080 task-api