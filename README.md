# Project Name

Brief project description.

## Table of Contents

- [Description](#description)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Configuration](#configuration)

## Description

This project allow to manage firmawares and devices, you can create, list, update and delete firmwares and devices. In order to do this you need to log in into the application with your username and password, if you do not have an account you can use the signup endpoint to register into the application.

## Getting Started

Explain how to get your project up and running. Include information about prerequisites and installation steps.

### Prerequisites

- Go (version 1.20)
- Gin (version 1.9.0)
- Postgres

### Installation

```bash
# Clone the repository
git clone https://github.com/dgsaltarin/FirmawaresManager.git

# Install dependencies
go mod tidy

# Build and run the project
go run main.go
```

## Configuration

In order to make the project work, you need to configure a postgres database, and put the connection configurations into a .env file, you can see how put those connection variable with the .env.example file

you can deploy de postgrest database instance using a docker image with:

- docker pull postgres
- docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=firmwares -d postgres
