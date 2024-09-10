# go-microservices

## Frontend

- Navigate to frontend service
  - `cd frontEnd/src/web`
- Run the frontend app
  - `go run main.go`

NOTE: Only if its first time (downloads the source code of the packages used and its dependencies within your workspace)

- `go install .src/web` or `go get .src/web`

## Broker service

- Navigate to Broker service
  - `cd brokerService/src/api`
- Run the broker service
  - `go run main.go routes.go handlers.go`
