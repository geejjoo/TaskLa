# TaskLa
Storage

Project for creating and canceling orders from a warehouse

## Installation

Clone the repository:

```bash
git clone https://github.com/geejjoo/TaskLa.git
```

Navigate to the cloned directory:
```bash
cd ./test_task/Lamoda_task
```

Install dependencies:
```bash
go mod tidy
```

## Usage
### Running the Project
To run the project, execute the following command in your terminal:
```bash
docker compose up
```
By default, the project will be available at `http://localhost:8000`.


## Project Structure
- `cmd/`: Directory for main executable files.
- `internal/`: Directory for packages used in the application.
  - `handler/`: Package containing HTTP request handlers.
  - `repository/`: Package containing the implementation of database operations.
  - `service/`: Package containing the business logic of the application.
- `configs/`: Directory for configuration files.
- `schema/`: Directory for database initialization.

# Retrieving reservation information
curl -i -X POST http://127.0.0.1:8000/api/v1/reserve/

# Response (successful request):
# HTTP/1.1 200 OK

# Response (unsuccessful request):
# HTTP/1.1 400 Bad Request

# Cancelling a reservation
curl -i -X POST http://127.0.0.1:8000/api/v1/reserve/cancel

# Response (successful request):
# HTTP/1.1 200 OK

# Response (unsuccessful request):
# HTTP/1.1 400 Bad Request

# Retrieving inventory information for a specific storage
curl -i -X GET http://127.0.0.1:8000/api/v1/inventory/Storage1

# Response (successful request):
# HTTP/1.1 200 OK

# Response (unsuccessful request):
# HTTP/1.1 404 Not Found


