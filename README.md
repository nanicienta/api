# My Dev Platform

Monorepo containing microservices built in Go, following Hexagonal Architecture principles.

## Overview

This repository provides the base structure and boilerplate for a set of microservices designed for developers. Each microservice follows hexagonal architecture, ensuring clear separation of concerns, testability, and maintainability.

### Services Included:

- **Account Service (`account-svc`)**: Handles user management, authentication, authorization, and organization management.
- **Application Service (`app-svc`)**: Manages applications and their respective environments.
- **Scaffold Service (`scaffold-svc`)**: Provides boilerplate generation for different technology stacks.
- **Logging Service (`logging-svc`)**: Centralized logging and monitoring service.

### Shared Libraries (`pkg`):

- **Logger**: Basic logging utilities.
- **Auth**: Authentication middleware.
- **Config**: Utilities for environment configuration loading.

---

## Prerequisites

- **Go 1.23+**
- ```go install mvdan.cc/gofumpt@latest```
- **Docker and Docker Compose** (optional for local setup)
- **Make** (optional for automation tasks)

---

## Getting Started

### 1. Clone the repository

```bash
git clone <your-repo-url>
cd my-dev-platform
```

### 2. Setup environment variables

Copy `.env.example` to `.env` and customize as needed:

```bash
cp .env.example .env
```

Edit `.env` to match your local configuration:

```env
DB_URL=postgres://user:pass@localhost:5432/dbname
```

### 3. Install dependencies

```bash
go mod tidy
```

---

## Running Services Locally

### Individual Service

Navigate to the service directory and run:

```bash
cd services/account-svc/cmd
go run account.go
```

Repeat this step for any other service you wish to run.

### Using Docker Compose (Recommended)

*(Docker Compose file not provided in this template, but highly recommended)*

A basic example `docker-compose.yml` structure:

```yaml
version: '3.8'
services:
  account_svc:
    build: ./services/account-svc
    environment:
      - DB_URL=${DB_URL}
    ports:
      - "9000:9000"

  app_svc:
    build: services/application-svc
    environment:
      - DB_URL=${DB_URL}
    ports:
      - "9001:9001"

  scaffold_svc:
    build: ./services/scaffold-svc
    ports:
      - "9002:9002"

  logging_svc:
    build: ./services/logging-svc
    ports:
      - "9003:9003"

  db:
    image: postgres:14
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
      - POSTGRES_DB=dbname
    ports:
      - "5432:5432"
```

Run all services:

```bash
docker-compose up
```

---

## Development and Formatting

The project includes a `.prettierrc` for consistent formatting. To format your files:

```bash
npm install -g prettier
prettier --write .
```

---

## Contribution

- Fork the repository.
- Create your feature branch (`git checkout -b feature/my-new-feature`).
- Commit your changes (`git commit -am 'Add some feature'`).
- Push to the branch (`git push origin feature/my-new-feature`).
- Create a new Pull Request.

---

## License

This project is open source and available under the [MIT License](LICENSE.md).