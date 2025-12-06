# Calendar API

A RESTful API service for calendar management built with Go, Echo framework, and PostgreSQL.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Database Setup](#database-setup)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [License](#license)

## Features

- User authentication (Sign In/Sign Out)
- Company management
- Property management with pricing
- Event scheduling and timeline
- Product management
- Photo storage and retrieval
- WebSocket support for real-time updates
- JWT-based authentication
- Swagger API documentation
- Database migration tools

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Echo v4
- **Database**: PostgreSQL (with GORM ORM)
- **Authentication**: JWT (JSON Web Tokens)
- **API Documentation**: Swagger UI
- **Configuration**: godotenv for environment variables
- **Validation**: go-playground/validator
- **CLI**: Cobra commands

## Prerequisites

- Go 1.24+
- PostgreSQL 12+
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/jihanlugas/calendar.git
cd calendar
```

2. Install dependencies:
```bash
go mod tidy
```

## Configuration

Create a `.env` file based on the `.env.example`:

```bash
cp .env.example .env
```

Then configure the following environment variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `DEBUG` | Debug mode | `true` |
| `SERVER_PORT` | Server port | `1323` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USERNAME` | Database username | `postgres` |
| `DB_PASSWORD` | Database password | `password` |
| `DB_NAME` | Database name | `calendar_db` |
| `CRYPTO_KEY` | Encryption key | `your_crypto_key` |
| `JWT_SECRET_KEY` | JWT secret key | `your_jwt_secret` |

## Database Setup

The application provides CLI commands for database management:

1. Run database migrations:
```bash
go run main.go db up
```

2. Seed the database with initial data:
```bash
go run main.go db seed
```

3. Reset the database (drop and recreate):
```bash
go run main.go db reset
```

4. Revert migrations:
```bash
go run main.go db down
```

## Running the Application

Start the server:
```bash
go run main.go
```

Or build and run:
```bash
go build -o calendar .
./calendar
```

The server will start on `http://localhost:1323` (or your configured port).

## API Documentation

When running in debug mode, you can access the Swagger UI documentation at:
```
http://localhost:1323/swg/
```

### Authentication Flow

1. Sign in to get authentication tokens:
```
POST /auth/sign-in
```

2. Use the token in the `Authorization` header for protected endpoints:
```
Authorization: Bearer <your_token_here>
```

## Project Structure

```
.
├── app/                 # Application modules
│   ├── auth/            # Authentication module
│   ├── company/         # Company management
│   ├── event/           # Event management
│   ├── photo/           # Photo handling
│   ├── product/         # Product management
│   ├── property/        # Property management
│   ├── propertygroup/   # Property group management
│   ├── propertyprice/   # Property pricing
│   ├── propertytimeline/# Property timeline
│   ├── user/            # User management
│   ├── usercompany/     # User-company relations
│   └── websocket/       # WebSocket handlers
├── cmd/                 # CLI commands
├── config/              # Configuration loader
├── constant/            # Constants definitions
├── cryption/            # Cryptography utilities
├── db/                  # Database connection
├── docs/                # Swagger documentation
├── jwt/                 # JWT utilities
├── model/               # Data models
├── request/             # Request DTOs
├── response/            # Response utilities
├── router/              # HTTP routes
└── utils/               # Utility functions
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.