# Hospital Management System

## Overview
The Hospital Management System is a robust backend API application built with Golang that provides comprehensive patient management and user authentication capabilities. Designed for healthcare environments, it supports role-based access control for receptionists and doctors with complete CRUD operations for patient records.

## Features
- **JWT Authentication**: Secure token-based authentication system
- **Role-Based Access Control**: Separate permissions for receptionists and doctors
- **Patient Management API**: Complete CRUD operations for patient records
- **User Management**: User registration and profile management
- **Middleware Integration**: Authentication, CORS, and logging middleware
- **PostgreSQL Integration**: Robust database layer with migrations
- **Comprehensive Testing**: Unit tests for all major components
- **RESTful API Design**: Clean and well-documented API endpoints

## Technology Stack
- **Backend**: Golang 1.20+ with Gin framework
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT tokens with bcrypt password hashing
- **Testing**: Testify framework with comprehensive test coverage
- **Architecture**: Clean architecture with repository pattern
- **Middleware**: Custom authentication, CORS, and logging

## Architecture
The application follows clean architecture principles:
- **Domain Layer**: Core business entities and interfaces
- **Service Layer**: Business logic and use cases
- **Infrastructure Layer**: Database repositories and external services
- **API Layer**: HTTP handlers, middleware, and routing
- **Utilities**: JWT, password hashing, and validation helpers

## Directory Structure
```
hospital-management-system/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/               # HTTP request handlers
│   │   ├── middleware/             # Authentication & CORS middleware
│   │   └── routes/                 # Route definitions
│   ├── config/                     # Configuration management
│   ├── domain/
│   │   ├── models/                 # Domain entities
│   │   └── repository/             # Repository interfaces
│   ├── infrastructure/
│   │   ├── database/               # Database connection & migrations
│   │   └── repository/             # Repository implementations
│   └── services/                   # Business logic layer
├── pkg/
│   └── utils/                      # JWT, hashing, validation utilities
├── tests/                          # Comprehensive test suite
│   ├── handlers/                   # Handler tests
│   ├── services/                   # Service layer tests
│   ├── middleware/                 # Middleware tests
│   ├── utils/                      # Utility tests
│   └── testutils/                  # Test helpers and mocks
├── web/                            # Static assets and templates
├── docs/                           # API documentation
├── .env.example                    # Environment variables template
├── docker-compose.yml              # Docker services configuration
├── Dockerfile                      # Container build configuration
├── Makefile                        # Build and test automation
├── go.mod                          # Go module dependencies
└── README.md                       # Project documentation
```

## API Endpoints

### Authentication
- `POST /api/auth/login` - User authentication
- `POST /api/auth/register` - User registration
- `POST /api/logout` - User logout (protected)

### Patient Management
- `GET /api/patients` - Get all patients (protected)
- `POST /api/patients` - Create new patient (protected)
- `GET /api/patients/:id` - Get patient by ID (protected)
- `PUT /api/patients/:id` - Update patient (protected)
- `DELETE /api/patients/:id` - Delete patient (protected)

### User Management
- `GET /api/users/:id` - Get user by ID (protected)
- `PUT /api/users/:id` - Update user (protected)

## Setup Instructions

### Prerequisites
- Go 1.20 or higher
- PostgreSQL 13+
- Make (optional, for using Makefile commands)

### Installation

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd hospital-management-system
   ```

2. **Install Dependencies**:
   ```bash
   make deps
   # or
   go mod download
   ```

3. **Environment Configuration**:
   ```bash
   cp .env.example .env
   ```
   Configure your `.env` file with:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=hospital_management
   JWT_SECRET=your_jwt_secret_key
   PORT=8080
   ```

4. **Database Setup**:
   ```bash
   # Create database
   createdb hospital_management
   
   # Run migrations
   psql -U <username> -d hospital_management -f internal/infrastructure/database/migrations/001_create_users_table.sql
   psql -U <username> -d hospital_management -f internal/infrastructure/database/migrations/002_create_patients_table.sql
   ```

5. **Build and Run**:
   ```bash
   make run
   # or
   go run cmd/server/main.go
   ```

### Docker Setup

1. **Using Docker Compose**:
   ```bash
   docker-compose up -d
   ```

2. **Build Docker Image**:
   ```bash
   make docker-build
   make docker-run
   ```

## Testing

The project includes comprehensive unit tests covering all major components:

### Run All Tests
```bash
make test
```

### Run Specific Test Categories
```bash
# Unit tests only
make test-unit

# Integration tests
make test-integration

# Generate coverage report
make test-coverage
```

### Test Structure
- **Handler Tests**: API endpoint testing with mocked dependencies
- **Service Tests**: Business logic testing with repository mocks
- **Middleware Tests**: Authentication and authorization testing
- **Utility Tests**: JWT, hashing, and validation testing

## Development

### Available Make Commands
```bash
make help              # Show all available commands
make build             # Build the application
make run               # Build and run the application
make test              # Run all tests
make test-coverage     # Generate test coverage report
make clean             # Clean build artifacts
make lint              # Run code linter
make fmt               # Format code
make vet               # Run go vet
make check             # Run all quality checks
```

### Code Quality
The project maintains high code quality through:
- **Linting**: golangci-lint integration
- **Testing**: Comprehensive unit test coverage
- **Code Formatting**: Automated go fmt
- **Static Analysis**: go vet integration

## Database Schema

### Users Table
- `id` (Primary Key)
- `username` (Unique)
- `password` (Hashed)
- `role` (receptionist/doctor)
- `created_at`, `updated_at`

### Patients Table
- `id` (Primary Key)
- `first_name`, `last_name`
- `date_of_birth`
- `gender`
- `phone_number`, `email`
- `address`
- `created_at`, `updated_at`

## Security Features
- **JWT Authentication**: Secure token-based authentication
- **Password Hashing**: bcrypt for secure password storage
- **Role-Based Access**: Different permissions for receptionists and doctors
- **Input Validation**: Comprehensive request validation
- **CORS Protection**: Configurable CORS middleware

## API Documentation
Comprehensive API documentation is available in the Postman collection:
[Hospital Management API Collection](https://www.postman.com/workspace/My-Workspace~ec53562f-b09d-4deb-851e-c2a71a34a647/collection/19612596-70dc12c8-2b99-4ac4-b724-0a889de7283c?action=share&creator=19612596&active-environment=19612596-c1b274f7-9779-474a-8542-632cdacbb186)