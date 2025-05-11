# Task Manager - Full Stack Application

A full-stack task management application built with Golang backend and React frontend.

## Project Structure

```
.
├── backend/           # Golang backend application
│   ├── cmd/          # Application entry points
│   ├── internal/     # Private application code
│   ├── pkg/          # Public library code
│   └── migrations/   # Database migrations
├── frontend/         # React frontend application
└── README.md         # This file
```

## Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- PostgreSQL 14 or later
- Make (optional, for using Makefile commands)

## Setup Instructions

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:(For PostgresDB)

4. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm start
   ```

## API Endpoints

- `GET /api/tasks` - Get all tasks
- `POST /api/tasks` - Create a new task
- `PUT /api/tasks/{id}` - Update a task
- `DELETE /api/tasks/{id}` - Delete a task

## Features

- Create, read, update, and delete tasks
- Task status management (Pending, In-Progress, Completed)
- Due date tracking
- Responsive UI
- Error handling and validation
- Database persistence

## Development

The project uses:
- Backend: Golang with Gin framework
- Frontend: React with TypeScript
- Database: PostgreSQL
- ORM: GORM 
