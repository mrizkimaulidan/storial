# storial

Storial.co API clone using Go. With minimal third party library, without using Gorm and Gin. This project shows you how to build idiomatic API in Go only with their built-in packages. Not all API endpoint are implemented in this project.

Library/frameworks used:
- MySQL Driver (https://github.com/go-sql-driver/mysql)
- JWT for API authorization (https://github.com/golang-jwt/jwt)
- Gorilla Mux for routing (https://github.com/gorilla/mux)
- Crypto for password hashing (https://github.com/golang/crypto)
- Godotenv (https://github.com/joho/godotenv)
- Ozzo-Validation (https://github.com/go-ozzo/ozzo-validation)

API Documentation with Postman: https://documenter.getpostman.com/view/10904143/VUxPv7MD

This project is not recommended to be use in production, because lack of security, code best practices. This project intended for learning and purposes only.

This project use Repository and Service pattern. And currently only support MySQL as the database.

Project installation:

Setup your database environment in .env file based on your credentials. Copy and rename `.env.example` file to `.env`.

Clone:
```bash
$ git clone https://github.com/mrizkimaulidan/storial.git
```

```bash
$ cd storial
```

Download required dependencies:
```bash
$ go mod download
```

Build:
```bash
$ go build -o cmd/main cmd/main.go
```

Run the compiled file:
```bash
$ cmd/main
```