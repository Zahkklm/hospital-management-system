# Hospital Management System

## Overview
The Hospital Management System is a web application designed to facilitate the management of patient records and user authentication for receptionists and doctors. The application provides a simple interface for receptionists to register and manage patients, while doctors can view and update patient information.

## Features
- **User Authentication**: Single login page for both receptionists and doctors.
- **Patient Management**: Receptionists can perform CRUD operations on patient records.
- **Doctor Access**: Doctors can view and update patient-related details.
- **Middleware**: Includes authentication, CORS, and logging middleware.
- **Database Integration**: Utilizes PostgreSQL for data storage with migration support.

## Technology Stack
- **Backend**: Golang with Gin framework
- **Database**: PostgreSQL
- **Frontend**: HTML, CSS, JavaScript (optional)
- **Testing**: Unit tests for handlers, services, and repositories

## Directory Structure
```
hospital-management-system
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers
│   │   ├── middleware
│   │   └── routes
│   ├── config
│   ├── domain
│   ├── infrastructure
│   └── services
├── pkg
│   └── utils
├── tests
├── web
│   ├── static
│   └── templates
├── docs
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Setup Instructions
1. **Clone the Repository**:
   ```
   git clone <repository-url>
   cd hospital-management-system
   ```

2. **Environment Variables**:
   Copy the `.env.example` to `.env` and configure your database connection settings.

3. **Database Migration**:
   Run the migration scripts to set up the database:
   ```
   psql -U <username> -d <database> -f internal/infrastructure/database/migrations/001_create_users_table.sql
   psql -U <username> -d <database> -f internal/infrastructure/database/migrations/002_create_patients_table.sql
   ```

4. **Run the Application**:
   ```
   go run cmd/server/main.go
   ```

5. **Access the Application**:
   Open your browser and navigate to `http://localhost:8080` to access the login page.

## API Documentation
API endpoints are documented in `docs/api.yaml`. You can use tools like Swagger or Postman to explore the API.

## Testing
Run the unit tests using:
```
go test ./tests
```

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

## Acknowledgments
- Inspired by various open-source hospital management systems.
- Thanks to the contributors and the community for their support.