# ManageLife

My personal task management (aka "Adulting") application.

## Development Notes

- `go run .`
- Rebuild = ctrl + shft + p; usually only when change .devcontainer files
- `go mod tidy`
```js
fetch('http://localhost:4000/api/tasks', {
  method: 'POST',
  body: JSON.stringify({
    title: "chrome title",
    notes: "chrome notes"
  }),
  headers: {
    'Content-type': 'application/json; charset=UTF-8'
  }
})
.then(res => res.json())
.then(console.log)
```

## Packages Using

- [Chi Router](https://github.com/go-chi/chi)
- go ent
  - `go generate ./..`
  - `go run -mod=mod entgo.io/ent/cmd/ent new Task`
  - `go generate ./ent`

## Notes to save and reference later in the project 

- `dateOnly := time.Date(2023, time.October, 7, 0, 0, 0, 0, time.UTC)`

## API Routs

base = `http://localhost:4000/api/tasks`
get a list of tasks = `http://localhost:4000/api/tasks/`

## Notes
```go
package main

func main() {
	var x int = 3

	z := &x // & = address of (creates a pointer to the location in memory where x is)
	zz := *z // * = dereference (undoes a pointer to a memory location returning the original memory at the location)

	var y *int // pointer to an int
}

```

## Troubleshooting

- container terminal not as expected: ctrl + shft + p --> rebuild