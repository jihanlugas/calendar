# Qwen Code Configuration for Go Calendar Project

## Project Information
- Project Name: Calendar Application
- Language: Go (Golang)
- Frameworks/Libraries: Standard library, Gin (if used), GORM (if used)
- Architecture Style: Modular (based on folders like app, config, model, router)

## Preferences
- Preferred Language for Communication: Indonesian
- Code Style: Follow Go idioms and clean architecture principles
- Testing Approach: Unit tests per package, integration tests for API endpoints
- Documentation Style: Inline comments + README.md updates

## Environment Setup Notes
- Working Directory: /home/jihanlugas/development/go/calendar
- Go Version: (to be checked in go.mod)
- Important Files:
  - Main entry point: main.go
  - Configuration files: config/
  - Database models: model/
  - API routes: router/
  - Request/response structs: request/, response/

## Common Tasks
- Run application: `go run main.go`
- Run tests: `go test ./...`
- Build application: `go build -o calendar-app`
- Check dependencies: `go mod tidy`

## Custom Notes
- Use `.env` for environment variables (example in `.env.example`)
- JWT handling modules are located in jwt/
- Cryptography utilities are in cryption/
- Constants are defined in constant/