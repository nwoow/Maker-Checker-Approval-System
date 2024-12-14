# Maker-Checker Approval System

A **Golang-based Maker-Checker Approval System** where users can send messages requiring approval from other users. The system tracks approvals and ensures messages are only sent after obtaining the required number of approvals. This project is built with **Golang**, **SQLite**, and containerized using **Docker** and **Docker Compose** for seamless deployment.

---

## Features

- **Send Messages**: Submit messages with a recipient and required approvers.
- **Approval Workflow**: Approvers validate messages, and only fully approved messages are marked as "Approved."
- **SQLite Persistence**: Persistent storage of messages using SQLite.
- **Containerized Deployment**: Simplified setup and deployment using Docker and Docker Compose.

---

## Technologies Used

- **Programming Language**: Golang
- **Database**: SQLite
- **Containerization**: Docker, Docker Compose
- **API Testing**: cURL/Postman
- **Frameworks/Libraries**:
  - `github.com/mattn/go-sqlite3`: SQLite driver
  - `github.com/google/uuid`: UUID generation

---

## Project Structure

```plaintext
maker_checker_sqlite/
├── Dockerfile                # Docker configuration for containerization
├── docker-compose.yml        # Docker Compose configuration
├── main.go                   # Entry point of the application
├── go.mod                    # Go module dependencies
├── go.sum                    # Dependency checksums
├── src/
│   ├── application/          # Application logic
│   ├── domain/               # Domain entities and interfaces
│   │   ├── entities/         # Message entity definition
│   │   ├── repositories/     # Repository interface
│   ├── infrastructure/       # API and persistence logic
│   │   ├── api/              # API handlers
│   │   └── persistence/      # SQLite repository implementation
│   ├── services/             # Business services (if needed)
└── data/                     # Directory for SQLite database


### Steps to Run the Application

#### 1. Clone the Repository:

```bash
git clone <repository_url>
cd maker_checker_sqlite


2. Build and Run with Docker Compose:
Build and start the application:

bash
Copy code
docker-compose up --build