# Rapatin Aja Backend

Backend service for Rapatin Aja application built with Go and MongoDB.

## Getting Started with Docker

You can run the entire stack (Go backend and MongoDB) using Docker Compose.

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. Clone the repository
2. Run the following command:

```bash
docker-compose up --build
```

The backend will be available at `http://localhost:8080`.

## Configuration

The application uses the following environment variables:

- `PORT`: The port on which the server will run (default: 8080)
- `MONGO_URL`: MongoDB connection string

In the Docker Compose setup, these are pre-configured to use the local MongoDB container.

## API Endpoints

### Auth
- `POST /api/register`: Register a new user
- `POST /api/login`: Login user

### Meetings
- `GET /api/meetings`: Get all meetings
- `GET /api/meetings/:id`: Get a meeting by ID
- `POST /api/meetings`: Create a new meeting
- `PATCH /api/meetings/:id`: Update a meeting
- `DELETE /api/meetings/:id`: Delete a meeting

### User
- `GET /api/user/:id`: Get user profile
- `PUT /api/user/:id`: Update user profile