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

1. Clone the Repository:

    git clone <repository_url>
    cd maker_checker_sqlite


2. Build and Run with Docker Compose:
   Build and start the application:

    docker-compose up --build

3. Verify the Application:

    curl http://localhost:9090  



## API Endpoints

1. Send a Message
Endpoint: POST /send
Description: Creates a new message.

Example Request:

    curl -X POST http://localhost:9090/send \
    -H "Content-Type: application/json" \
    -d '{"content":"Hello, please approve this!", "recipient":"user@example.com", "max_approvers":2}'

Example Response:

    {
    "id": "generated-message-id"
    }


2. Approve a Message
Endpoint: POST /approve
Description: Approves the message with the given ID.

Example Request:

    curl -X POST http://localhost:9090/approve \
    -H "Content-Type: application/json" \
    -d '{"message_id":"<message_id>", "approver":"approver1"}'

Replace <message_id> with the id from the /send API response.

Example Response:

    HTTP/1.1 200 OK

If the message is already approved or not in a pending state:

    {
    "error": "message is not in pending state"
    }


## Logs and Debugging
   1. View Logs
   To view container logs:

    docker-compose logs -f