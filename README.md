# Code Assist API

A simple Go Fiber backend service that provides code execution simulation, auto-fixing, and help features.

## Features

- **Run Code**: Simulate code execution with basic syntax checking
- **Auto-Fix**: Automatically fix common code formatting issues
- **Help**: Get help with common programming concepts

## API Endpoints

### 1. Run Code
```bash
POST /run
Content-Type: application/json

{
    "code": "package main\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}"
}
```

### 2. Auto-Fix Code
```bash
POST /autofix
Content-Type: application/json

{
    "code": "package main\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")"
}
```

### 3. Get Help
```bash
POST /help
Content-Type: application/json

{
    "query": "function"
}
```

## Running the Server

1. Install dependencies:
```bash
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:3001`

## Testing with cURL

### Test Run Endpoint
```bash
curl -X POST http://localhost:3001/run \
  -H "Content-Type: application/json" \
  -d '{"code":"package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello\")\n}"}' | jq
```

### Test Auto-Fix Endpoint
```bash
curl -X POST http://localhost:3001/autofix \
  -H "Content-Type: application/json" \
  -d '{"code":"package main\n\nfunc main() {\n    fmt.Println(\"Hello\")' | jq
```

### Test Help Endpoint
```bash
curl -X POST http://localhost:3001/help \
  -H "Content-Type: application/json" \
  -d '{"query": "function"}' | jq
```

## Implementation Details

- Uses Go Fiber for the web framework
- Simple in-memory processing (no database required)
- Rule-based code fixing (no AI/ML)
- Basic syntax checking for common errors
- Help system with keyword matching
