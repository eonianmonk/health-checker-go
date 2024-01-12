# health-checker-go
Health checker for web-services

To run do 
```bash
go run ./cmd run
```
For "feature for stop ping process if one of url respond with fail" add flag --stopIfFailed
To set port for service set arg --port. Important: add ":" before port number 
To set ping timeout use arg --timeout
For example:
```bash
go run ./cmd run --stopIfFailed --port=":8008" --timeout=3
```

To run all tests run:
```bash
go test ./...
```
