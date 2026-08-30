# Go URL Shortener

A lightweight URL shortening service built with **Go**, designed with a clean backend architecture and separation between the HTTP server, application logic, and storage layer.

The project focuses on building a simple backend service while keeping the codebase structured and easy to extend.

## ✨ Features

* 🔗 Create short URLs from long URLs
* ↪️ Redirect short URLs to their original destinations
* ⚡ Lightweight Go HTTP server
* 🧩 Separation of application and storage logic
* 📦 Clean project structure
* 🚀 Designed to be easy to extend

---

## 🏗️ Architecture

The project follows a layered structure that separates the server, internal application logic, and persistence.

```text
                     ┌───────────────┐
                     │     Client    │
                     └───────┬───────┘
                             │
                             ▼
                     ┌───────────────┐
                     │  HTTP Server  │
                     │ cmd/server    │
                     └───────┬───────┘
                             │
                             ▼
                     ┌───────────────┐
                     │    Internal   │
                     │ Application   │
                     │    Logic      │
                     └───────┬───────┘
                             │
                             ▼
                     ┌───────────────┐
                     │    Storage    │
                     │     Layer     │
                     └───────────────┘
```

This separation makes it possible to change the persistence implementation without tightly coupling it to the HTTP layer.

---

## 📁 Project Structure

```text
go-url-shortner/
│
├── cmd/
│   └── server/
│       └── ...
│
├── internal/
│   └── ...
│
├── storage/
│   └── ...
│
├── go.mod
├── go.sum
└── .gitignore
```

### `cmd/server`

Contains the application entry point and server setup.

### `internal`

Contains the application's internal business logic and components that aren't intended to be imported by external packages.

### `storage`

Contains persistence-related functionality and keeps storage concerns separated from the rest of the application.

---

## 🔄 How It Works

### Creating a Short URL

```text
Client
  │
  │ Long URL
  ▼
HTTP Server
  │
  ▼
Application Logic
  │
  ▼
Generate Short ID
  │
  ▼
Storage
  │
  ▼
Short URL
```

### Redirecting

```text
Client
  │
  │ /abc123
  ▼
HTTP Server
  │
  ▼
Application Logic
  │
  ▼
Storage Lookup
  │
  ▼
Original URL
  │
  ▼
HTTP Redirect
```

---

## 🧰 Tech Stack

* **Go**
* **Go HTTP server**
* **Go modules**
* Layered application architecture
* Pluggable storage abstraction

---

## 🚀 Getting Started

### Prerequisites

Install:

* [Go](https://go.dev/)

### Clone the repository

```bash
git clone https://github.com/shaurya21343/go-url-shortner.git

cd go-url-shortner
```

### Install dependencies

```bash
go mod download
```

### Run the server

```bash
go run ./cmd/server
```

The server will start using the configuration defined by the application.

---

## 🧪 Development

Build the application:

```bash
go build ./...
```

Run the project's tests:

```bash
go test ./...
```

Run tests with additional output:

```bash
go test -v ./...
```

---

## 🎯 Project Goals

This project was built to explore backend development with Go and understand how to structure a small web service beyond putting everything into a single file.

The main focus areas are:

* Go project organization
* HTTP server development
* Separation of concerns
* Storage abstractions
* Backend API design
* Writing maintainable Go code

---

## 🔮 Future Improvements

Potential improvements include:

* [ ] Custom aliases
* [ ] URL expiration
* [ ] Click analytics
* [ ] Rate limiting
* [ ] Persistent production database
* [ ] Docker support
* [ ] API documentation
* [ ] Automated tests
* [ ] Structured logging
* [ ] Metrics and observability
* [ ] Authentication
* [ ] Horizontal scaling

---

## 👨‍💻 Author

**Shaurya Singh**

GitHub:
https://github.com/shaurya21343

---

## 📄 License

This project is open source. See the repository for licensing information.
