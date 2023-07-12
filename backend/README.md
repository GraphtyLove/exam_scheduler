# Mock exam scheduler backend
This is the backend for the mock exam scheduler. It is a simple REST API that allows the frontend to create, read, update and delete exams.

## Technologies
- [go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [MongoDB](https://www.mongodb.com/)

## Requirements
Create a `.env` file in the `backend` folder. Add the connection string to your MongoDB database in the `.env` file. The connection string should look like this:
```
MONGO_CONNECTION_STRING=mongodb+srv://<USER_HERE>:<PASSWORD_HERE>@<DB_URL_HERE>/?retryWrites=true&w=majority
```

## Running the backend
### Dev mode (hot reload)
Install [air](https://github.com/cosmtrek/air)
```bash
go install github.com/cosmtrek/air@latest
```

Run the backend
```bash
air
```

### Production mode
Build the backend
```bash
go build
```
Run the backend
```bash
./backend
```

## Run the tests
```bash
go test ./...
```