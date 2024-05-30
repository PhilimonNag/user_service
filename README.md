# User gRPC Service

This is a gRPC service for managing user details with search functionality.

## Build and Run

### Prerequisites

- Go 1.21.3
- Docker

### Running Locally

1. Clone the repository

```
git clone github.com/philimonnag/user_service
cd user_service
```

2. Run the application

```
go mod tidy
go run .
```

3. Run The unit test

```
go test -v
```

4. Example grpcurl for interact

Get user by ID:

```
grpcurl -d '{"id": 1}' -plaintext localhost:50051 user.UserService/GetUserById
```

Get users by IDs:

```
grpcurl -d '{"ids": [1, 2]}' -plaintext localhost:50051 user.UserService/GetUsersByIds
```

Search users by city:

```
grpcurl -d '{"city": "LA"}' -plaintext localhost:50051 user.UserService/SearchUsers
```

Create a new user:

```
grpcurl -d '{"fname": "Alice", "city": "SF", "phone": 1122334455, "height": 5.7, "married": false}' -plaintext localhost:50051 user.UserService/CreateUser
```
