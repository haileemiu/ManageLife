# ManageLife

My personal task management (aka "Adulting") application.

## Development Notes

- `go run .`
- Rebuild = ctrl + shft + p; usually only when change .devcontainer files
- `go mod tidy`
- `http://localhost:4000/api/tasks/` = list

## Packages Using

- https://github.com/go-chi/chi
- go ent
  - `go generate ./..`
  - `go run -mod=mod entgo.io/ent/cmd/ent new Task`
  - `go generate ./ent`

## Notes to save and reference later in the project 
- `dateOnly := time.Date(2023, time.October, 7, 0, 0, 0, 0, time.UTC)`
