# Nanicienta

Monorepo containing microservices built in Go, trying to follow Hexagonal Architecture principles.

## Overview
This repository provides the base structure and boilerplate for a set of microservices designed for developers. Each microservice follows hexagonal architecture, ensuring clear separation of concerns, testability, and maintainability.

### Services Included:

- **Account Service (`[account-svc](services/account-svc)account-svc`)**: Handles user management.
- **Application Service (`[application-svc](services/application-svc)`)**: Manages applications and their respective configuration. Like names, where is going to be deployed, etc.
- **Authentication Service (`[authentication-svc](services/authentication-svc)`)**: Manages authentication.
- **Authorization Service (`[authorization-svc](services/authorization-svc)`)**: Manages authorization.
- **Billing Service (`[billing-svc](services/billing-svc)`)**: Provides billing and payment functionalities.
- **Environment Service (`[environment-svc](services/environment-svc)`)**: Provides environemnt functionalities. like env vars, ephemeral environments.
- **Logging Service (`logging-svc`)**: Centralized logging and monitoring service.
- **Namespaces Service (`[namespaces-svc](services/namespaces-svc)`)**: Manages namespaces, access to the components, and other functionalities.
- **Organization Service (`[organization-svc](services/organization-svc)`)**: Handles the creation and management of organizations. Authentication prefrences
- **Scaffold Service (`[scaffold-svc](services/scaffold-svc)`)**: In the future this is going to be responsible for scaffolding new services in different languages, and other functionalities.

### Shared Libraries (`pkg`):
- **[command](pkg/domain/command)**: Some common commands used by the services.
- **[errors](pkg/domain/errors)**: Error utils and error codes.
- **[auth](pkg/domain/auth)**: Authentication middleware.
- **[logging](pkg/ports/logging)**: Logging port
- **[client](pkg/ports/outbound/client)**: Clients used by the services

---

## Prerequisites

- **Go 1.24+**
- ```go install mvdan.cc/gofumpt@latest```
- **Docker and Docker Compose** (optional for local setup)
- **Make** (optional for automation tasks)

---

## Getting Started

### 1. Clone the repository

```bash
git clone git@github.com:nanicienta/api.git
cd api
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

## Run all services:

```bash
docker-compose up
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