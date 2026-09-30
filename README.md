# Auth Microservice

A production-ready microservice responsible for user authentication, authorization.

## 🚀 Features

- **User Authentication:** Sign Up, Sign In.
- **Token Management:** JWT.
- **Security:** Password hashing.

## 🛠️ Tech Stack

- **Language/Framework:** Golang, gRPC
- **Database:** PostgreSQL 

## 📋 Prerequisites

Before running this service, ensure you have the following installed:
- Docker and Docker Compose
- Golang 1.27.1 +

## ⚙️ Getting Started

### 1. Clone the repository 
```bash
git clone https://github.com/AndroDeMohawk/sso-app
cd sso-app
```

### 2. Environment Variables
Create a `.env` file in the root directory and configure your variables:
```env
CONFIG_PATH="./config/local.yaml" #optional
```

### 3. Installation & Running

```bash
# Install dependencies
go mod download

#Run app:
task run

#testing:
task test

#migrations:
task migrate:up
task migrate:down
```

**Using Docker:**
```bash
docker-compose up --build
docker-compose dowm -v
```

## 🔌 API Endpoints
```
#for most comfortable using endpoints you can download https://github.com/AndroDeMohawk/protos
#and in Insomnia/Postman select a .proto file by directory "protos/proto/sso/sso.proto"
```
| Method | Endpoint | Description |
| :--- | :--- | :--- | 
| `gRPC` | `http://localhost:44044/Auth/Register` | Register a new user | 
| `gRPC` | `http://localhost:44044/Auth/Login` | Authenticate user & get tokens | 
| `gRPC` | `http://localhost:44044/Auth/IsAdmin` | Checking user role | 
