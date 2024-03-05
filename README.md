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

# Получение информации о резервировании
curl -i http://127.0.0.1:8000/api/v1/reserve/

## Ответ (успешный запрос):
## HTTP/1.1 200 OK

## Ответ (неуспешный запрос):
## HTTP/1.1 400 Bad Request

## Отмена резервирования
curl -i -X POST http://127.0.0.1:8000/api/v1/reserve/cancel

## Ответ (успешный запрос):
## HTTP/1.1 200 OK

## Ответ (неуспешный запрос):
## HTTP/1.1 400 Bad Request

## Получение информации о запасах в определенном складе
curl -i http://127.0.0.1:8000/api/v1/inventory/Storage1

## Ответ (успешный запрос):
## HTTP/1.1 200 OK

## Ответ (неуспешный запрос):
## HTTP/1.1 404 Not Found


