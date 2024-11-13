# Hydrate

[![Go Version][go-image]][go-url]
[![Gin Version][gin-image]][gin-url]
[![License][license-image]][license-url]

A robust and scalable RESTful API built with Go (Golang) for tracking daily water intake. This API allows users to register, authenticate, record water intake events, and retrieve their intake history.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Database Setup](#database-setup)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [Water Intake](#water-intake)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [License](#license)

---

## Features

- **User Registration and Authentication**: Secure user registration and JWT-based authentication.
- **Water Intake Recording**: Users can log their water intake events with timestamps.
- **History Retrieval**: Retrieve water intake records with optional date filters.
- **Input Validation**: Robust input validation using Gin binding and validation tags.
- **Secure Password Storage**: Passwords are hashed using bcrypt.
- **Modular Architecture**: Clean separation of concerns with services, repositories, and handlers.
- **Configuration Management**: Easy configuration using Viper.
- **Logging**: Structured logging with Logrus.
- **Scalable Design**: Ready for containerization and deployment.

---

## Architecture

The application follows a layered architecture:

- **Model**: Defines the data structures.
- **Repository**: Handles database interactions.
- **Service**: Contains business logic.
- **Handler**: Processes HTTP requests and responses.
- **Middleware**: Contains authentication logic.
- **Router**: Defines API routes and middleware.

---

## Getting Started

### Prerequisites

- **Go**: Version 1.23 or higher
- **PostgreSQL**: Version 17 or higher
- **Git**: For version control
- **Docker** (Optional): For containerization

### Installation

Clone the repository:

```bash
git clone https://github.com/pp00x/hydrate.git
cd hydrate
```

Install dependencies:

```bash
go mod tidy
```

### Configuration

The application uses [Viper](https://github.com/spf13/viper) for configuration management. You can configure the application using environment variables or by creating a `config.yaml` file in the root directory.

#### Environment Variables

Set the following environment variables:

```bash
export SERVER_PORT=8080
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_DBNAME=hydrate_db
export DATABASE_PASSWORD=your_db_password
export JWT_SECRET_KEY=your_secret_key
```

#### `config.yaml` File

Alternatively, create a `config.yaml` file:

```yaml
server:
  port: "8080"
  read_timeout: 5s
  write_timeout: 5s

database:
  host: "localhost"
  port: "5432"
  user: "postgres"
  dbname: "water_tracker"
  password: "your_db_password"
  sslmode: "disable"
  timezone: "UTC"

jwt:
  secret_key: "your_secret_key"
```

### Database Setup

Ensure PostgreSQL is installed and running. Create the database:

```sql
CREATE DATABASE water_tracker;
```

The application will automatically run database migrations on startup.

### Running the Application

Run the application:

```bash
go run cmd/server/main.go
```

The server will start on the configured port (default is `8080`).

---

## API Endpoints

### Authentication

#### Register a New User

- **Endpoint**: `POST /api/v1/register`
- **Request Body**:

  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  ```json
  {
    "message": "Registration successful"
  }
  ```

#### Login

- **Endpoint**: `POST /api/v1/login`
- **Request Body**:

  ```json
  {
    "email": "john@example.com",
    "password": "your_password"
  }
  ```

- **Response**:

  ```json
  {
    "token": "your_jwt_token"
  }
  ```

### Water Intake

#### Record Water Intake

- **Endpoint**: `POST /api/v1/water-intake`
- **Headers**:

  ```
  Authorization: Bearer your_jwt_token
  ```

- **Request Body**:

  ```json
  {
    "amount": 250,
    "taken_at": "2021-10-01T08:30:00Z"
  }
  ```

- **Response**:

  ```json
  {
    "message": "Water intake recorded"
  }
  ```

#### Get Water Intake Records

- **Endpoint**: `GET /api/v1/water-intake`
- **Headers**:

  ```
  Authorization: Bearer your_jwt_token
  ```

- **Query Parameters (Optional)**:

  - `start_date`: `YYYY-MM-DD`
  - `end_date`: `YYYY-MM-DD`

- **Example**:

  ```
  GET /api/v1/water-intake?start_date=2021-09-01&end_date=2021-10-01
  ```

- **Response**:

  ```json
  {
    "data": [
      {
        "id": 1,
        "user_id": 1,
        "amount": 250,
        "taken_at": "2021-10-01T08:30:00Z",
        "created_at": "2021-10-01T08:30:00Z",
        "updated_at": "2021-10-01T08:30:00Z"
      }
    ]
  }
  ```

---

## Project Structure

```
water-intake-tracker/
├── cmd/
│   └── server/        # Entry point of the application
├── config/            # Configuration files and initialization
├── internal/          # Application modules and business logic
│   ├── handler/       # HTTP handlers
│   ├── middleware/    # Authentication middleware
│   ├── model/         # Database models
│   ├── repository/    # Database interactions
│   ├── router/        # API routes
│   └── service/       # Business logic and services
├── pkg/               # Shared utilities (logging, database)
├── go.mod             # Go module file
├── go.sum             # Go dependencies checksum file
└── README.md          # Project documentation
```

---

## Technologies Used

- **Go**
- **Gin Web Framework**
- **GORM**
- **PostgreSQL**
- **JWT**
- **Viper**
- **Logrus**
- **bcrypt**

## License

Licensed under the Apache-2.0 License.

[go-image]: https://img.shields.io/badge/Go-1.23.3-blue.svg
[go-url]: https://golang.org/doc/go1.23
[gin-image]: https://img.shields.io/badge/gin--gonic-1.10.0-blue
[gin-url]: https://github.com/gin-gonic/gin
[license-image]: https://img.shields.io/badge/license-Apache-blue.svg
[license-url]: LICENSE
