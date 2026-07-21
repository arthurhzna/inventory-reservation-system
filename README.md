# inventory-reservation-system
## Running Tests

The project includes unit tests for the inventory reservation business logic.

Current unit tests are located at:

```text
backend/internal/application/usecase/inventory_usecase_test.go
```

To run the inventory use case tests:

```bash
cd backend
go test -v ./internal/application/usecase
```

To run all tests:

```bash
cd backend
go test -v ./...
```

To execute all tests with Go race detection enabled:

```bash
cd backend
go test -race -v ./...
```

The race detector verifies that the application does not contain data races during concurrent execution.