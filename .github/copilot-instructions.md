# GitHub Copilot Instructions

## General Rules

- Use Go (Golang) as the primary programming language.
- Follow idiomatic Go conventions.
- Always run `gofmt` formatting.
- Prefer clear and readable code over clever optimizations.
- Avoid unnecessary abstractions.
- Use English for all code, comments, commit messages, and variable names.

---

## Project Structure

This project follows layered / clean architecture.

Example structure:

```txt
/app
  /user
    handler.go
    usecase.go
    repository.go
/pkg
/config
/utils
```

### Layer Responsibilities

#### handler

Responsibilities:
- HTTP request/response handling
- validation
- parsing query/body/path params

#### usecase

Responsibilities:
- business logic
- orchestration
- transaction flow

#### repository

Responsibilities:
- database queries only
- no business logic

---

## Coding Style

### Function Rules

- Keep functions small and focused.
- Split large functions into smaller helper functions.
- Avoid deeply nested conditions.
- Return early whenever possible.

Bad example:

```go
if err == nil {
    if user != nil {
        ...
    }
}
```

Good example:

```go
if err != nil {
    return err
}

if user == nil {
    return nil
}
```

---


## Error Handling

- Always handle errors explicitly.
- Never ignore returned errors.
- Wrap errors with context when needed.

Bad example:

```go
err := repo.Create(ctx, data)
if err != nil {
    return err
}
```

```go
if err := repo.Create(ctx, data); err != nil {
    return err
}
```

Good example:

```go
err := repo.Create(ctx, data)
if err != nil {
    return fmt.Errorf("failed create user: %w", err)
}
```

### Avoid Inline Error Assignment

Do not use inline error assignment inside `if` statements.

Bad example:

```go
if err := myfunc(); err != nil {
	return err
}
```

Good example:

```go
err = myfunc()
if err != nil {
    return err
}
```

---

## Naming Convention

### Use clear naming

Good:

```go
CreateInstance
GetOrganizationByID
UpdateUserBalance
```

Bad:

```go
DoData
HandleStuff
Process
```

### Interface Naming

Use capability-based names.

Good:

```go
type UserRepository interface {}
type PaymentService interface {}
```

Avoid:

```go
type IUserRepository interface {}
```

---

## Security

- Never hardcode secrets.
- Use environment variables.
- Validate all user input.
- Sanitize query parameters when needed.

---

## Performance

- Avoid N+1 queries.
- Select only required fields.
- Use pagination for list endpoints.
- Avoid loading unnecessary relations.

---

## Refactoring Rules

When refactoring:
- Do not change existing behavior.
- Do not rename variables unless necessary.
- Preserve API contract.
- Split long functions into smaller functions safely.

---

## Copilot Behavior

When generating code:
- Prioritize readability and maintainability.
- Follow existing project patterns.
- Reuse existing helper/util functions if available.
- Avoid introducing new dependencies unless necessary.
- Generate production-ready code whenever possible.

When generating SQL or GORM queries:
- Always include proper filtering.
- Consider soft delete behavior.
- Avoid full table scans if possible.

When generating APIs:
- Include validation.
- Include error handling.
- Return consistent response format.
